// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define the route for GET /v1/test/empty_get
	router.GET("/v1/test/empty_get", func(c *gin.Context) {
		// Respond with an empty JSON response (just to simulate a fast response)
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	// Run the server on port 8083
	router.Run(":8083")
}
