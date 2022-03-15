package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/hello", Hello)

	port := ":1234"

	fmt.Printf("http://localhost%s\n", port)

	err := router.Run(port)
	if err != nil {
		panic(err)
	}

	// /:name c.Param("name")
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "howdy")
}

// Marshal

// Middleware
