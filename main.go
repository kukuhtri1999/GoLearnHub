// main.go
package main

import (
	"fmt"
	"net/http"
	"os"

	_ "golearnhub/docs"
	"golearnhub/models"
	"golearnhub/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // Import swaggerFiles
	ginSwagger "github.com/swaggo/gin-swagger" // Import gin-swagger middleware
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// @title GoLearnHub API
// @description API for GoLearnHub online learning platform
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	err := loadEnv()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	initDB()

	router := gin.Default()

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register user routes
	routes.RegisterUserRoutes(router, db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoLearnHub!"})
	})

	fmt.Println("Server is running on :8080")
	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func initDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Database connected successfully!")

	// Auto Migrate the database models
	db.AutoMigrate(&models.User{}, &models.Course{}, &models.Lecture{}, &models.Quiz{}, &models.UserCourseProgress{})
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}