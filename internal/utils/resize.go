package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"math"
)

func Resize(uploadFile *UploadFile, width, height int) ([]byte, image.Image) {

	fmt.Println("uploadFile.IMG.Bounds()", uploadFile.IMG.Bounds().Max.X, uploadFile.IMG.Bounds().Max.Y)

	dstW, dstH := width, height
	srcH := uploadFile.IMG.Bounds().Dy()
	srcW := uploadFile.IMG.Bounds().Dx()

	if uploadFile.IMG.Bounds().Dx() < width {
		width = uploadFile.IMG.Bounds().Dx()
	}

	if dstW == 0 {
		tmpW := float64(dstH) * float64(srcW) / float64(srcH)
		dstW = int(math.Max(1.0, math.Floor(tmpW+0.5)))
	}
	if dstH == 0 {
		tmpH := float64(dstW) * float64(srcH) / float64(srcW)
		dstH = int(math.Max(1.0, math.Floor(tmpH+0.5)))
	}

	// Resize srcImage to width = 800px preserving the aspect ratio.
	newImg := imaging.Fit(uploadFile.IMG, dstW, dstH, imaging.Lanczos)

	fmt.Println("resize", newImg.Bounds())

	var byff bytes.Buffer
	if uploadFile.Format == "jpeg" {
		jpeg.Encode(&byff, newImg, nil)
	}

	if uploadFile.Format == "png" {
		png.Encode(&byff, newImg)
	}

	return byff.Bytes(), newImg

}

// reverseOrientation amply`s what ever operation is necessary to transform given orientation
// to the orientation 1
func reverseOrientation(img image.Image, o string) *image.NRGBA {
	switch o {
	case "1":
		return imaging.Clone(img)
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}

	return imaging.Clone(img)
}
