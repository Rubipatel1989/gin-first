package handlers

import (
	"net/http"
	"strconv"

	"mobile-api-service/database"
	"mobile-api-service/models"

	"github.com/gin-gonic/gin"
)

// API for Frontend - Get User List
func GetUserList(c *gin.Context) {
	var users []models.User
	
	// Get query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Get only active users
	result := database.DB.Where("status = ?", "active").
		Offset(offset).
		Limit(limit).
		Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var total int64
	database.DB.Model(&models.User{}).Where("status = ?", "active").Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

