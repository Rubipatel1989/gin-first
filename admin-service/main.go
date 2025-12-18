package main

import (
	"log"

	"admin-service/admin"
	"admin-service/config"
	"admin-service/database"

	"github.com/GoAdminGroup/go-admin/engine"
	adminConfig "github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Connect to MySQL
	database.ConnectMySQL()

	// Connect to Redis (optional for admin panel)
	database.ConnectRedis()

	// Create Gin router
	r := gin.Default()

	// Initialize GoAdmin
	eng := engine.Default()
	
	// Configure GoAdmin
	adminCfg := &adminConfig.Config{
		Databases: adminConfig.DatabaseList{
			"default": {
				Host:   config.AppConfig.DBHost,
				Port:   config.AppConfig.DBPort,
				User:   config.AppConfig.DBUser,
				Pwd:    config.AppConfig.DBPassword,
				Name:   config.AppConfig.DBName,
				Driver: "mysql",
			},
		},
		UrlPrefix: "admin",
		Store: adminConfig.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    "en",
		Debug:       config.AppConfig.GinMode == "debug",
		ColorScheme: "skin-black",
		Title:       "Admin Panel",
		Logo:        "GoAdmin",
		MiniLogo:    "GA",
	}

	// Setup GoAdmin plugins
	if err := admin.SetupGoAdmin(eng); err != nil {
		log.Fatal("Failed to setup GoAdmin:", err)
	}

	// Add GoAdmin to Gin router
	if err := eng.AddConfig(adminCfg).Use(r); err != nil {
		log.Fatal("Failed to initialize GoAdmin:", err)
	}

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "admin-panel",
		})
	})

	// Start server
	serverPort := ":" + config.AppConfig.ServerPort
	log.Printf("Admin Panel starting on port %s", config.AppConfig.ServerPort)
	log.Printf("GoAdmin panel available at http://localhost%s/admin", serverPort)
	if err := r.Run(serverPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

