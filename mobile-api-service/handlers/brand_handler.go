package handlers

import (
	"net/http"
	"strconv"

	"mobile-api-service/database"
	"mobile-api-service/models"

	"github.com/gin-gonic/gin"
)

// API for Frontend - Get Brand List
func GetBrandList(c *gin.Context) {
	var brands []models.Brand
	
	// Get query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Get only active brands
	result := database.DB.Where("status = ?", "active").
		Offset(offset).
		Limit(limit).
		Find(&brands)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brands"})
		return
	}

	var total int64
	database.DB.Model(&models.Brand{}).Where("status = ?", "active").Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    brands,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

