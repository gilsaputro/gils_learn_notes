package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Middleware to print path and method
	router.Use(func(c *gin.Context) {
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Endpoint with request JSON
	router.POST("/api/json", func(c *gin.Context) {
		var jsonBody map[string]interface{}
		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, jsonBody)
	})

	// Endpoint with a dynamic parameter {id}
	router.GET("/api/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	// Endpoint with query parameter
	router.GET("/api/query", func(c *gin.Context) {
		queryParam := c.Query("param")
		c.JSON(http.StatusOK, gin.H{"query_param": queryParam})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
