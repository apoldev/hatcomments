package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	AttachmentImage    string = "image"
	AttachmentVideo    string = "video"
	AttachmentSticker  string = "sticker"
	AttachmentTenorGif string = "tenor_gif"
)

// gorm.Model definition
type Attachment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Hash    string `gorm:"index;type:varchar(255)" json:"hash"`
	Type    string `gorm:"type:varchar(255)" json:"type"`
	Alt     string `json:"alt"`
	Name    string `json:"name"`
	Preview string `json:"preview"`

	Format string `gorm:"-" json:"format"`

	UserID uint `gorm:"index" json:"user_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// gorm.Model definition
type AttachmentComment struct {
	ID uint `gorm:"primaryKey" json:"id"`

	CommentID    uint `json:"comment_id"`
	AttachmentID uint `json:"attachment_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
