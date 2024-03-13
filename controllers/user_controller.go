// controllers/user_controller.go
package controllers

import (
	"net/http"

	"golearnhub/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterUser handles user registration with additional validation and processing
// @Summary Register a new user
// @Accept json
// @Produce json
// @Param user body models.User true "User information for registration"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Email already exists"
// @Router /users/register [post]
func RegisterUser(c *gin.Context, db *gorm.DB) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email uniqueness
	var existingUser models.User
	result := db.Where("email = ?", newUser.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Validate password complexity (e.g., at least 8 characters, one uppercase, one digit)
	if len(newUser.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters"})
		return
	}

	if !containsUppercase(newUser.Password) || !containsDigit(newUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain at least one uppercase letter and one digit"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	// Update user object with hashed password
	newUser.Password = string(hashedPassword)

	// Create the user in the database
	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Helper function to check if a string contains an uppercase letter
func containsUppercase(s string) bool {
	for _, char := range s {
		if 'A' <= char && char <= 'Z' {
			return true
		}
	}
	return false
}

// Helper function to check if a string contains a digit
func containsDigit(s string) bool {
	for _, char := range s {
		if '0' <= char && char <= '9' {
			return true
		}
	}
	return false
}
