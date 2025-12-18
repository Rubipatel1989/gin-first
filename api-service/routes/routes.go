package routes

import (
	"api-service/handlers"
	"api-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORS())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "api-service",
			"message": "Server is running",
		})
	})

	// API Routes for Frontend (Flutter/Mobile/Web)
	api := r.Group("/api")
	{
		// User endpoints
		api.GET("/users", handlers.GetUserList)

		// Brand endpoints
		api.GET("/brands", handlers.GetBrandList)

		// Store endpoints
		api.GET("/stores", handlers.GetStoreList)
	}

	return r
}

