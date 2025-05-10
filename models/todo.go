package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Text      string         `json:"text"`
	Done      bool           `json:"done"`
	UserID uint   `json:"userId"`  
	User      *User          `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}