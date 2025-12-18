#!/bin/bash

echo "Installing GoAdmin dependencies..."

# Install main GoAdmin package
go get github.com/GoAdminGroup/go-admin@latest

# Install GoAdmin Gin adapter (this is part of go-admin, but we need to ensure it's available)
go get github.com/GoAdminGroup/go-admin/adapter/gin@latest

# Install GoAdmin MySQL driver (this is part of go-admin)
go get github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql@latest

# Install GoAdmin themes
go get github.com/GoAdminGroup/themes@latest

# Install other dependencies
go get github.com/gin-gonic/gin@latest
go get github.com/go-redis/redis/v8@latest
go get github.com/joho/godotenv@latest
go get gorm.io/driver/mysql@latest
go get gorm.io/gorm@latest

# Clean up and verify
go mod tidy

echo "Dependencies installed successfully!"
echo "Run 'go run main.go' to start the server"

