package main

import (
	"log"

	"mobile-api-service/config"
	"mobile-api-service/database"
	"mobile-api-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Connect to MySQL
	database.ConnectMySQL()

	// Connect to Redis
	database.ConnectRedis()

	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	serverPort := ":" + config.AppConfig.ServerPort
	log.Printf("API Service starting on port %s", config.AppConfig.ServerPort)
	log.Printf("API available at http://localhost%s/api", serverPort)
	if err := r.Run(serverPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

