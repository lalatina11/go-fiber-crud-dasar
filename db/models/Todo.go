package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string         `json:"title" gorm:"not null;unique"`
	Description string         `json:"description" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
