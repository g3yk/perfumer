package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseGormModel struct {
	ID        uint       `json:"id" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2025-04-23T00:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2025-04-23T00:00:00Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"2025-04-23T00:00:00Z"`
}

type BaseUser struct {
	Username string `json:"username" example:"user2"`
	Email    string `json:"email" example:"user2@me.com"`
	Password string `json:"password" example:"user1234"`
	Name     string `json:"name" example:"John Doe"`
}

type UserResponse struct {
	BaseGormModel
	BaseUser
}

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:50;" validate:"required,min=3,max=50" json:"username"`
	Email    string `gorm:"uniqueIndex;not null;size:255;" validate:"required,email" json:"email"`
	Password string `gorm:"not null;" validate:"required,min=6,max=50" json:"password"`
	Name     string `json:"name"`
}
