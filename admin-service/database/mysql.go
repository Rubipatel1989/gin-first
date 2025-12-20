package database

import (
	"fmt"
	"log"

	"admin-service/config"
	"admin-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to MySQL database:", err)
	}

	log.Println("MySQL database connected successfully")

	// Auto migrate tables
	err = DB.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Brand{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}

// InitGoAdminTables creates the necessary GoAdmin system tables
func InitGoAdminTables() {
	// Create goadmin_session table
	DB.Exec("CREATE TABLE IF NOT EXISTS goadmin_session (" +
		"id varchar(150) NOT NULL PRIMARY KEY, " +
		"`values` text NOT NULL, " +
		"created_at timestamp NULL DEFAULT NULL, " +
		"updated_at timestamp NULL DEFAULT NULL" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")

	// Create goadmin_users table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_users (
			id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
			username varchar(150) NOT NULL,
			password varchar(150) NOT NULL,
			name varchar(150) NOT NULL,
			avatar varchar(255) DEFAULT NULL,
			remember_token varchar(100) DEFAULT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			UNIQUE KEY username (username)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_roles table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_roles (
			id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name varchar(50) NOT NULL,
			slug varchar(50) NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			UNIQUE KEY slug (slug)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_permissions table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_permissions (
			id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name varchar(50) NOT NULL,
			slug varchar(50) NOT NULL,
			http_method varchar(255) DEFAULT NULL,
			http_path text NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			UNIQUE KEY slug (slug)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_menu table
	DB.Exec("CREATE TABLE IF NOT EXISTS goadmin_menu (" +
		"id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY, " +
		"parent_id int unsigned DEFAULT 0, " +
		"type tinyint unsigned DEFAULT 0, " +
		"`order` int unsigned DEFAULT 0, " +
		"title varchar(50) NOT NULL, " +
		"icon varchar(50) NOT NULL, " +
		"uri varchar(50) DEFAULT NULL, " +
		"header varchar(150) DEFAULT NULL, " +
		"created_at timestamp NULL DEFAULT NULL, " +
		"updated_at timestamp NULL DEFAULT NULL" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")

	// Create goadmin_operation_log table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_operation_log (
			id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
			user_id int unsigned NOT NULL,
			path varchar(255) NOT NULL,
			method varchar(10) NOT NULL,
			ip varchar(15) NOT NULL,
			input text NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			KEY user_id (user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_user_permissions table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_user_permissions (
			user_id int unsigned NOT NULL,
			permission_id int unsigned NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			PRIMARY KEY (user_id, permission_id),
			KEY permission_id (permission_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_role_users table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_role_users (
			role_id int unsigned NOT NULL,
			user_id int unsigned NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			PRIMARY KEY (role_id, user_id),
			KEY user_id (user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_role_permissions table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_role_permissions (
			role_id int unsigned NOT NULL,
			permission_id int unsigned NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			PRIMARY KEY (role_id, permission_id),
			KEY permission_id (permission_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_role_menu table
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS goadmin_role_menu (
			role_id int unsigned NOT NULL,
			menu_id int unsigned NOT NULL,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			PRIMARY KEY (role_id, menu_id),
			KEY menu_id (menu_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	`)

	// Create goadmin_site table
	DB.Exec("CREATE TABLE IF NOT EXISTS goadmin_site (" +
		"id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY, " +
		"`key` varchar(100) NOT NULL, " +
		"`value` text, " +
		"description varchar(3000) DEFAULT NULL, " +
		"type int unsigned DEFAULT 0, " +
		"state tinyint unsigned DEFAULT 0, " +
		"created_at timestamp NULL DEFAULT NULL, " +
		"updated_at timestamp NULL DEFAULT NULL, " +
		"UNIQUE KEY `key` (`key`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	
	// Add state column if table exists but column is missing
	var columnExists int
	DB.Raw("SELECT COUNT(*) FROM information_schema.columns WHERE table_schema = ? AND table_name = 'goadmin_site' AND column_name = 'state'", config.AppConfig.DBName).Scan(&columnExists)
	if columnExists == 0 {
		DB.Exec("ALTER TABLE goadmin_site ADD COLUMN state tinyint unsigned DEFAULT 0 AFTER type")
	}

	// Insert default admin user if it doesn't exist
	var count int64
	DB.Table("goadmin_users").Where("username = ?", "admin").Count(&count)
	if count == 0 {
		// Default password is "admin" (hashed with bcrypt)
		// GoAdmin default password hash for "admin" is: $2a$10$E9x8zJ8qJ8qJ8qJ8qJ8qJ.8qJ8qJ8qJ8qJ8qJ8qJ8qJ8qJ8qJ8qJ
		// But we'll let GoAdmin create it through the install page
		log.Println("GoAdmin system tables initialized. Please visit /admin/install to create the admin user.")
	} else {
		log.Println("GoAdmin system tables initialized. Admin user already exists.")
	}

	log.Println("GoAdmin system tables initialized successfully")
}

