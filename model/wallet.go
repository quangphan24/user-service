package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallets struct {
	BaseModel
	ID      string `json:"id"`
	Name    string `json:"name" validate:"required"`
	UserId  string `json:"user_id" validate:"required"`
	Balance int    `json:"balance"`
}

func (w *Wallets) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = uuid.New().String()
	return nil
}
