package main

import (
	"fmt"
	"os"

	"github.com/NeerChayaphon/CodaAssignment/SimpleAPI/game"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get the port from an environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a Gin router
	router := gin.Default()
	gameHandler := game.NewGameHandler()

	// Define a route to handle POST requests
	router.POST("/", gameHandler.Topup)

	// Start the backend server on the specified port
	err := router.Run(":" + port)
	if err != nil {
		fmt.Printf("Error starting Backend: %v\n", err)
	}
}
