// models/user_course_progress_model.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type UserCourseProgress struct {
	ID               uint   `gorm:"primaryKey" json:"progressID"`
	UserID           uint   `json:"userID"`
	CourseID         uint   `json:"courseID"`
	CompletionStatus string `json:"completionStatus"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (ucp *UserCourseProgress) BeforeCreate(tx *gorm.DB) (err error) {
	ucp.CreatedAt = time.Now()
	return
}

func (ucp *UserCourseProgress) BeforeUpdate(tx *gorm.DB) (err error) {
	ucp.UpdatedAt = time.Now()
	return
}
