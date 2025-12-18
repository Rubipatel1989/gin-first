package routes

import (
	"gin-first/handlers"
	"gin-first/middleware"

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
			"message": "Server is running",
		})
	})

	// Note: GoAdmin handles /admin routes for the admin panel
	// The old API routes are kept for backward compatibility but can be removed
	// GoAdmin provides full CRUD through its interface

	// API Routes for Flutter
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUserList)
		api.GET("/brands", handlers.GetBrandList)
	}

	return r
}

