package models

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
type Vote struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Vote      int  `json:"vote"`
	CommentID uint `json:"comment_id"`
	UserID    uint `json:"user_id"`

	User      *User          `json:"user,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
