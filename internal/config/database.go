// internal/config/database.go
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB creates a new GORM database connection
func NewDB() (*gorm.DB, error) {
    dsn := "user=postgres dbname=golearnhub password=postgres host=localhost port=5432 sslmode=disable"
    // Replace youruser, yourdb, and yourpassword with your actual PostgreSQL credentials

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
