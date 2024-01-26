package utils

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type UploadFile struct {
	Alt  string `json:"alt"`
	Hash string `json:"hash"`

	Directory string `json:"directory"`
	File      []byte `json:"file"`
	FileWebp  []byte `json:"file_webp"`

	Type   string `json:"type"`
	Uuid   string `json:"uuid"`
	Format string `json:"format"`

	IMG      image.Image `json:"-"`
	SmallIMG image.Image `json:"-"`
}

// image formats and magic numbers
var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

// mimeFromIncipit returns the mime type of an image file from its first few
// bytes or the empty string if the file does not look like a known file type
func MimeFromIncipit(incipit []byte) string {
	incipitStr := string(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(incipitStr, magic) {
			return mime
		}
	}

	return ""
}

func ReadUploadFile(data *multipart.FileHeader) (*UploadFile, error) {

	file, err := data.Open()
	defer file.Close()

	if err != nil {
		return nil, err
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	h := sha1.New()
	if _, err := h.Write(fileData); err != nil {
		return nil, err
	}

	fileFormat := ""
	fileFormatSplit := strings.Split(data.Filename, ".")
	if len(fileFormatSplit) > 1 {
		fileFormat = fileFormatSplit[1]
	}

	fileType := "image"

	contentType := MimeFromIncipit(fileData)
	if contentType == "" {
		contentType = data.Header.Get("Content-Type")
	}

	contentTypePart := strings.Split(contentType, "/")
	if len(contentTypePart) == 2 {
		fileFormat = contentTypePart[1]
	}

	if contentType == "" || fileFormat == "" {
		return nil, errors.New("not image")
	}

	//fmt.Println("MimeFromIncipit", MimeFromIncipit(fileData))

	result := UploadFile{
		Alt:  data.Filename,
		File: fileData,
		Hash: fmt.Sprintf("%x", h.Sum(nil)),
		Type: fileType,

		Uuid:   uuid.NewString(),
		Format: fileFormat,

		Directory: getDateFormat(time.Now()),
	}

	return &result, nil

}

func WriteUploadFile(file *UploadFile) (*UploadFile, error) {

	// запись
	dirDate := "./storage/upload/" + file.Directory + "/"

	reader := bytes.NewReader(file.File)
	var img image.Image
	var err error

	if file.Format == "jpeg" {
		img, err = jpeg.Decode(reader)
		if err != nil {
			return nil, err
		}
	}

	if file.Format == "png" {
		img, err = png.Decode(reader)
		if err != nil {
			return nil, err
		}
	}

	reader = bytes.NewReader(file.File)
	x, err := exif.Decode(reader)

	if err != nil {
		fmt.Println(err)
	}

	if x != nil {
		orient, _ := x.Get(exif.Orientation)
		if orient != nil {
			//fmt.Println("%s had orientation %s", orient.String())
			img = reverseOrientation(img, orient.String())
		} else {
			//fmt.Println("%s had no orientation - implying 1")
			img = reverseOrientation(img, "1")
		}

	}

	file.IMG = img

	// create DIR
	if _, err := os.Stat(dirDate); os.IsNotExist(err) {

		var dirMod uint64
		if dirMod, err = strconv.ParseUint("0777", 8, 32); err == nil {
			err = os.Mkdir(dirDate, os.FileMode(dirMod))

			if err != nil {
				fmt.Println("err", err)
			}

		}
	}

	if file.Format == "gif" {
		// Гифку положим как есть
		err = os.WriteFile(dirDate+file.Uuid+"."+file.Format, file.File, 0644)
		if err != nil {
			return nil, fmt.Errorf("Couldn't write out file: %s", err)
		}

	} else {

		// Изменим размер перед записью, ограничив ширину 1920px
		bigImageBytes, _ := Resize(file, 1920, 0)
		err = os.WriteFile(dirDate+file.Uuid+"."+file.Format, bigImageBytes, 0644)
		if err != nil {
			return nil, fmt.Errorf("Couldn't write out file: %s", err)
		}

		smallImageBytes, _ := Resize(file, 0, 100)
		err = os.WriteFile(dirDate+file.Uuid+"."+file.Format+"_small."+file.Format, smallImageBytes, 0644)
		if err != nil {
			return nil, fmt.Errorf("Couldn't write out file: %s", err)
		}

	}

	// Закодируем в webp, если возможно
	//EncodeToWebp(file)
	//
	//if len(file.FileWebp) > 0 {
	//	path := path + ".webp"
	//	err := os.WriteFile(path, file.FileWebp, 0644)
	//	if err != nil {
	//		return nil, fmt.Errorf("Couldn't write out file: %s", err)
	//	}
	//}

	return file, nil

}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

func EncodeToWebp(uploadFile *UploadFile) {

	if uploadFile.IMG == nil {
		return
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	var buff bytes.Buffer

	if err := webp.Encode(&buff, uploadFile.IMG, options); err != nil {
		log.Fatalln(err)
	}

	uploadFile.FileWebp = buff.Bytes()

	fmt.Println("Success webp")

}
