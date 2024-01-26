package models

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Type     string `gorm:"type:varchar(255)" json:"type"`
	SocialID string `json:"social_id"`

	Email string `json:"email"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`

	Role   string `gorm:"type:varchar(255)" json:"role"`
	Icon   string `gorm:"type:varchar(255)" json:"icon"`
	Status string `gorm:"type:varchar(255)" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Comments           []Comment           `json:"comments,omitempty"`
	UserAuthorizations []UserAuthorization `json:"user_authorizations,omitempty"`
	UserBans           []UserBan           `json:"user_bans,omitempty"`
}

type UserAuthorization struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `json:"user_id"`

	IP string `gorm:"type:varchar(255)" json:"ip"`
	UA string `json:"ua"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `json:"-"`
}
