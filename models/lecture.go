// models/lecture_model.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Lecture struct {
	ID         uint   `gorm:"primaryKey" json:"lectureID"`
	CourseID   uint   `json:"courseID"`
	LectureTitle string `json:"lectureTitle"`
	LectureType  string `json:"lectureType"`
	LectureURL   string `json:"lectureURL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (l *Lecture) BeforeCreate(tx *gorm.DB) (err error) {
	l.CreatedAt = time.Now()
	return
}

func (l *Lecture) BeforeUpdate(tx *gorm.DB) (err error) {
	l.UpdatedAt = time.Now()
	return
}
