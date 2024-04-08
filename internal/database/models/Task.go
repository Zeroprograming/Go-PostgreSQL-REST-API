package models

import "gorm.io/gorm"

// Task struct represents a task entity in the database.
type Task struct {
	gorm.Model // Embedded struct for adding standard fields (ID, CreatedAt, UpdatedAt, DeletedAt)

	Title       string `gorm:"type:varchar(100);not null"` // Title of the task
	Description string // Description of the task
	Done        bool   `gorm:"default:false"` // Indicates whether the task is done or not
	UserId      uint   // ID of the user associated with the task
}
