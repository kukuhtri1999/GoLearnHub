// models/quiz_model.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	ID        uint   `gorm:"primaryKey" json:"quizID"`
	CourseID  uint   `json:"courseID"`
	Question  string `json:"question"`
	Options   string `json:"options"`
	CorrectAnswer string `json:"correctAnswer"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	q.CreatedAt = time.Now()
	return
}

func (q *Quiz) BeforeUpdate(tx *gorm.DB) (err error) {
	q.UpdatedAt = time.Now()
	return
}
