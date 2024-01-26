package repo

import (
	"cloud_payments/internal/comment"
	"cloud_payments/internal/models"
	"cloud_payments/internal/utils"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"html"
	"strconv"
	"time"
)

type CommentRepository struct {
	Db *gorm.DB
}

func (r *CommentRepository) GetComments(limit, offset int) []models.Comment {

	var comments []models.Comment

	r.Db.Model(&models.Comment{}).
		Preload("User").
		Preload("Parent").
		Preload("Attachments").
		Preload("Votes").
		Preload("Votes.User").
		Order("id desc").
		Limit(limit).
		Offset(offset).
		Find(&comments)

	return comments
}

func (r *CommentRepository) GetProject() (project *models.Project) {
	return
}

func (r *CommentRepository) GetRoomByID(roomID int) *models.Room {

	room := models.Room{}

	tx := r.Db.Model(&models.Room{}).Where("id = ?", roomID).First(&room)

	if tx.Error != nil {
		return nil
	}

	return &room

}

func (r *CommentRepository) GetRoomOrCreate(projectID int, roomName string) *models.Room {

	room := models.Room{}

	tx := r.Db.Model(&models.Room{}).Where("name = ?", roomName).First(&room)

	hash := md5.Sum([]byte(roomName))

	if tx.Error != nil {
		room = models.Room{
			ProjectID: uint(projectID),
			Name:      roomName,
			Slug:      hex.EncodeToString(hash[:]),
		}

		r.Db.Create(&room)
	}

	return &room

}

func (r *CommentRepository) LoadCommentsByRoom(room *models.Room, limit, offset int) (result []models.Comment) {

	ids := []uint{}

	if offset == 0 {
		r.Db.Raw(`select c2.id from (
select * from comments where level is null and room_id = ? order by created_at desc limit ?) as c
left join comments c2 ON c2.path <@ c.path;`, room.ID, limit).Scan(&ids)
	} else {
		r.Db.Raw(`select c2.id from (
select * from comments where level is null and room_id = ? and id < ? order by created_at desc limit ?) as c
left join comments c2 ON c2.path <@ c.path;`, room.ID, offset, limit).Scan(&ids)
	}

	if len(ids) == 0 {
		return
	}

	var comments []models.Comment

	r.Db.Model(&models.Comment{}).Unscoped().
		Preload("User").
		Preload("Attachments").
		Preload("Votes").
		Order("created_at desc").
		Find(&comments, ids)

	// поиск родителей в выгрузке - 8 микросекунд
	for i := range comments {
		// Если это ответ, то найдем родителя
		if comments[i].ReplyTo != nil {
			for l := range comments {
				if comments[l].ID == *comments[i].ReplyTo {
					comments[i].Parent = &comments[l]
					break
				}
			}
		}
	}

	return comments
}

func (r *CommentRepository) GetCommentByID(id uint) *models.Comment {

	if id < 1 {
		return nil
	}

	c := models.Comment{}

	tx := r.Db.Model(&models.Comment{}).Unscoped().Where("id = ?", id).First(&c)

	fmt.Println(tx.Error, c)
	if tx.Error != nil {
		return nil
	}

	return &c

}

func (r *CommentRepository) EditComment(id uint, text string, attachments []int) (*models.Comment, error) {

	if text == "" && len(attachments) < 1 {
		return nil, errors.New("empty text or attachments")
	}

	c := r.GetCommentByID(id)

	if c == nil {
		return nil, errors.New("not found")
	}

	// Привяжем вложенные файлы
	if len(attachments) > 0 {

		attSlice := []models.Attachment{}
		r.Db.Model(&models.Attachment{}).Find(&attSlice, attachments)
		r.Db.Model(&c).Association("Attachments").Replace(attSlice)

	}

	c.Text = text

	tx := r.Db.Omit("Attachments").Save(&c)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return c, nil

}

func (r *CommentRepository) CreateComment(user models.User, room *models.Room, parent *models.Comment, text string, attachments []int) (*models.Comment, error) {

	if room.ID < 1 {
		return nil, errors.New("no room")
	}

	if text == "" && len(attachments) < 1 {
		return nil, errors.New("empty text or attachments")
	}

	if len(text) > 0 {
		text = html.EscapeString(text)
	}

	c := models.Comment{
		RoomID: room.ID,
		Text:   text,
		User:   user,
	}

	// Если есть родитель, тогда укажем уровень и ссылку на родителя
	if parent != nil {

		// Запретим ответ на сообщение в другой комнате
		if parent.RoomID != room.ID {
			return nil, errors.New("Incorrect room")
		}

		level := uint(1)
		if parent.Level != nil {
			level = *parent.Level + 1
		}

		c.ReplyTo = &parent.ID
		c.Level = &level
	}

	tx := r.Db.Create(&c)

	// Если создадлось, тогда обновим path
	if tx.RowsAffected > 0 {
		path := strconv.Itoa(int(c.ID))

		if parent != nil {
			path = parent.Path + "." + path
		}

		r.Db.Model(&c).Updates(models.Comment{Path: path})

		// Привяжем вложенные файлы
		if len(attachments) > 0 {

			attSlice := []models.Attachment{}

			r.Db.Model(&models.Attachment{}).Find(&attSlice, attachments)

			r.Db.Model(&c).Association("Attachments").Append(attSlice)
		}

	}

	return &c, nil

}

func (r *CommentRepository) GetTenorAttachment(data *comment.TenorGIFData) *models.Attachment {

	attach := models.Attachment{}

	tx := r.Db.Model(&models.Attachment{}).Where("hash = ?", data.ID).First(&attach)

	if tx.Error != nil {

		attach = models.Attachment{
			Name:    data.Gif,
			Preview: data.NanoGif,
			Hash:    data.ID,
			Alt:     data.Title,
			Type:    models.AttachmentTenorGif,
		}

		r.Db.Create(&attach)

	}

	if attach.ID > 0 {
		return &attach
	}

	return nil

}

func (r *CommentRepository) GetUploadedAttachment(file *utils.UploadFile) (*models.Attachment, bool) {

	attach := models.Attachment{}

	currentlyCreated := false

	fmt.Println("file.Hash", file.Hash)
	tx := r.Db.Model(&models.Attachment{}).Where("hash = ?", file.Hash).First(&attach)

	if tx.Error != nil {
		// Если не найдено - создадим

		name := file.Directory + "/" + file.Uuid + "." + file.Format
		preview := name + "_small." + file.Format

		if file.Format == "gif" {
			preview = name
		}

		attach = models.Attachment{
			Name:    name,
			Preview: preview,
			Hash:    file.Hash,
			Type:    file.Type,
			Alt:     file.Alt,
		}

		r.Db.Create(&attach)
		currentlyCreated = true
	}

	if attach.ID > 0 {
		return &attach, currentlyCreated
	}

	return nil, currentlyCreated

}

func (r *CommentRepository) GetOrCreateUserVoteOnComment(user *models.User, comment2 *models.Comment, vote int) (*models.Vote, error) {

	v := models.Vote{}

	r.Db.Model(&models.Vote{}).Where("user_id = ? and comment_id = ?", user.ID, comment2.ID).First(&v)

	// Если найден, то надо записать новый голос заменой значения vote
	if v.ID > 0 {
		r.Db.Model(&v).Update("vote", vote)
		return &v, nil
	}

	v = models.Vote{
		UserID:    user.ID,
		CommentID: comment2.ID,
		Vote:      vote,
	}

	tx := r.Db.Create(&v)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v, nil

}

func (r *CommentRepository) GetVotesByCommentID(id uint) (result []models.Vote, err error) {

	tx := r.Db.Model(&models.Vote{}).Where("comment_id = ?", id).Find(&result)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return

}

func (r *CommentRepository) DeleteComment(id int, u *models.User) error {

	tx := r.Db.Model(&models.Comment{
		ID: uint(id),
	}).Updates(map[string]interface{}{
		"deleted_by_id": u.ID,
		"deleted_at":    time.Now(),
	})

	return tx.Error

}

func (r *CommentRepository) RestoreComment(id int) error {

	tx := r.Db.Model(&models.Comment{
		ID: uint(id),
	}).Unscoped().Updates(map[string]interface{}{
		"deleted_at":    nil,
		"deleted_by_id": nil,
	})

	return tx.Error

}

func (r *CommentRepository) SaveCommentAction(action string, uid uint, old, newComment *models.Comment) error {

	ch := models.CommentHistory{
		Action:    action,
		CommentID: old.ID,
		UserID:    uid,
		Data:      "{}",
		NewData:   "{}",
	}

	if action == "edit" || newComment == nil && old != nil {
		bytes, _ := json.Marshal(&old)
		ch.Data = string(bytes)
	}

	if newComment != nil {

		var bytes []byte

		if action == "vote" && len(newComment.Votes) > 0 {
			bytes, _ = json.Marshal(&newComment.Votes[0])
		} else {
			bytes, _ = json.Marshal(&newComment)
		}

		ch.NewData = string(bytes)
	}

	r.Db.Model(&models.CommentHistory{}).Create(&ch)

	return nil
}
