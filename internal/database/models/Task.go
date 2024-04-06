package models

import "gorm.io/gorm"

// Task struct
type Task struct {
	gorm.Model

	Title       string `gorm:"type:varchar(100);not null"`
	Description string
	Done        bool `gorm:"default:false"`
	UserId      uint
}
