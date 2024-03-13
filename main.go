// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoLearnHub!"})
	})

	fmt.Println("Server is running on :8080")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
