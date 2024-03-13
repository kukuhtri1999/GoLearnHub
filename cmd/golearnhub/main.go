// cmd/golearnhub/main.go
package main

import "github.com/gin-gonic/gin"

func main() {
    // Initialize a GIN router
    router := gin.Default()

    // Define a simple route
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, GoLearnHub!",
        })
    })

    // Run the server on port 8080
    router.Run(":8080")
}
