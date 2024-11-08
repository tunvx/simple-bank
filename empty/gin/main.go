// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define the route for GET /v1/get_empty_connection
	router.GET("/v1/get_empty_connection", func(c *gin.Context) {
		// Respond with an empty JSON response (just to simulate a fast response)
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection is empty",
		})
	})

	// Run the server on port 8888
	router.Run(":8888")
}
