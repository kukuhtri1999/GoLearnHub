// migrations/migrate.go
package migrations

import (
	"golearnhub/models"

	"gorm.io/gorm"
)

// Migrate runs the database migrations
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Course{})
	db.AutoMigrate(&models.Lecture{})
	db.AutoMigrate(&models.Quiz{})
	db.AutoMigrate(&models.UserCourseProgress{})
}
