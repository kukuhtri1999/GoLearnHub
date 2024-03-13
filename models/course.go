// models/course_model.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID           uint   `gorm:"primaryKey" json:"courseID"`
	UserID       uint   `json:"userID"`
	CourseName   string `json:"courseName"`
	CourseDescription string `json:"courseDescription"`
	Subject      string `json:"subject"`
	Level        string `json:"level"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	return
}

func (c *Course) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
