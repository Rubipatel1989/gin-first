package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Address     string         `json:"address" gorm:"type:text"`
	Phone       string         `json:"phone" gorm:"type:varchar(50)"`
	Email       string         `json:"email" gorm:"type:varchar(255)"`
	Status      string         `json:"status" gorm:"type:varchar(50);default:active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateStoreRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" binding:"omitempty,email"`
}

type UpdateStoreRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" binding:"omitempty,email"`
	Status  string `json:"status"`
}

