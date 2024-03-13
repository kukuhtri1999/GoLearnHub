// routes/user_routes.go
package routes

import (
	"golearnhub/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", func(c *gin.Context) {
			controllers.RegisterUser(c, db)
		})
	}
}
