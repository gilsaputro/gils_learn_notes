package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware to print path and method
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Printf("Request: %s %s\n", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	})

	// Endpoint with request JSON
	e.POST("/api/json", func(c echo.Context) error {
		var jsonBody map[string]interface{}
		if err := c.Bind(&jsonBody); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, jsonBody)
	})

	// Endpoint with a dynamic parameter {id}
	e.GET("/api/id/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, map[string]string{"id": id})
	})

	// Endpoint with query parameter
	e.GET("/api/query", func(c echo.Context) error {
		queryParam := c.QueryParam("param")
		return c.JSON(http.StatusOK, map[string]string{"query_param": queryParam})
	})

	// Start the Echo server on port 8080
	e.Start(":8080")
}
