package handlers

import (
	"net/http"
	"strconv"

	"mobile-api-service/database"
	"mobile-api-service/models"

	"github.com/gin-gonic/gin"
)

// API for Frontend - Get Store List
func GetStoreList(c *gin.Context) {
	var stores []models.Store
	
	// Get query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Get only active stores
	result := database.DB.Where("status = ?", "active").
		Offset(offset).
		Limit(limit).
		Find(&stores)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stores"})
		return
	}

	var total int64
	database.DB.Model(&models.Store{}).Where("status = ?", "active").Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stores,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

