package models

import (
	"time"
)

// gorm.Model definition
type UserBan struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint   `json:"user_id"`
	IP     string `gorm:"type:varchar(255)" json:"ip"`

	BannedBy uint   `json:"banned_by"`
	Comment  string `json:"comment"`

	BannedUntil time.Time `json:"banned_until"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	User User `json:"-"`
}
