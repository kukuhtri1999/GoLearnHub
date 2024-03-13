// routes/user_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kukuhtri1999/GoLearnHub/controllers" // Update with your project's import path
)

// SetUserRoutes defines user-related routes
func SetUserRoutes(router *gin.Engine) {
	userController := controllers.UserController{}

	// Group user routes under "/user"
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("/profile", userController.GetUserProfile)
	}
}
