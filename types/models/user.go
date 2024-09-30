package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID      string         `gorm:"index" json:"user_id"`
	PhoneNumber string         `json:"phone_number"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Pin         string         `json:"pin"`
	Address     string         `json:"address"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
}

type UserRegisterPayload struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Pin         string `json:"pin"`
	Address     string `json:"address"`
}
