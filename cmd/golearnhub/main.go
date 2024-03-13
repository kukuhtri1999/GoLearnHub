// cmd/golearnhub/main.go
package main

import (
	"fmt"
	"golearnhub/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Initialize a GIN router
	router := gin.Default()

	// Initialize GORM database connection
	db, err := config.NewDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Migrate the database schema
	err = migrateDB(db)
	if err != nil {
		fmt.Println("Failed to migrate the database:", err)
		return
	}

	// Close the database connection when the application exits
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, GoLearnHub!",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}

// migrateDB performs database schema migration
func migrateDB(db *gorm.DB) error {
	// Add your database migration logic here
	// Example: db.AutoMigrate(&YourModel{})

	return nil
}
