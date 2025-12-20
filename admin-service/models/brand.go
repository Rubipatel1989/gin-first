package models

import (
	"time"

	"gorm.io/gorm"
)

type Brand struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text"`
	Logo        string         `json:"logo" gorm:"type:varchar(500)"`
	Status      string         `json:"status" gorm:"type:varchar(50);default:active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateBrandRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type UpdateBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Status      string `json:"status"`
}

