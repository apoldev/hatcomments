package models

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
type Comment struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	RoomID      uint `json:"room_id"`
	UserID      uint `json:"user_id"`
	DeletedByID uint `json:"deleted_by_id"` //  кем удалено

	Level   *uint  `gorm:"index" json:"level"`
	ReplyTo *uint  `json:"reply_to"`
	Text    string `json:"text"`

	Like int `gorm:"-" json:"like"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Path        string       `gorm:"type:ltree" json:"path"`
	User        User         `json:"user"`
	Attachments []Attachment `gorm:"many2many:attachment_comments;" json:"attachments"`
	Votes       []Vote       `json:"votes"`

	Parent *Comment `gorm:"-" json:"parent"`
}

// gorm.Model definition
type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Host      string         `gorm:"type:varchar(255)" json:"host"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Rooms []Room `json:"-"`
}

// gorm.Model definition
type Room struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ProjectID uint           `json:"project_id"`
	Name      string         `json:"name"`
	Slug      string         `gorm:"type:varchar(255)" json:"slug"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Project Project `json:"-"`
}

type CommentHistory struct {
	ID uint `gorm:"primaryKey" json:"id"`

	CommentID uint   `json:"comment_id"`
	UserID    uint   `json:"user_id"`
	Action    string `json:"action"`

	Data    string `gorm:"type:jsonb" json:"data"`
	NewData string `gorm:"type:jsonb" json:"newdata"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
