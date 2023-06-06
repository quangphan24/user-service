package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	ID       string `json:"id"`
	UserName string `json:"user_name" validate:"required" gorm:"not null"`
	Password string `json:"password" validate:"required" gorm:"not null"`
	Email    string `json:"email" validate:"required,email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	Id           string `json:"id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
