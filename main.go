package main

import (
	"log"

	"gin-first/admin"
	"gin-first/config"
	"gin-first/database"
	"gin-first/routes"

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

	// Connect to Redis
	database.ConnectRedis()

	// Setup routes
	r := routes.SetupRoutes()

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
	}

	// Setup GoAdmin plugins
	if err := admin.SetupGoAdmin(eng); err != nil {
		log.Fatal("Failed to setup GoAdmin:", err)
	}

	// Add GoAdmin to Gin router
	if err := eng.AddConfig(adminCfg).Use(r); err != nil {
		log.Fatal("Failed to initialize GoAdmin:", err)
	}

	// Start server
	serverPort := ":" + config.AppConfig.ServerPort
	log.Printf("Server starting on port %s", config.AppConfig.ServerPort)
	log.Printf("GoAdmin panel available at http://localhost%s/admin", serverPort)
	if err := r.Run(serverPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

