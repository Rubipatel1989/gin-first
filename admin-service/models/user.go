package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Phone     string         `json:"phone" gorm:"type:varchar(50)"`
	Status    string         `json:"status" gorm:"type:varchar(50);default:active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
}

type UpdateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email" binding:"omitempty,email"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

