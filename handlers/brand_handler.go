package handlers

import (
	"net/http"
	"strconv"

	"gin-first/database"
	"gin-first/models"

	"github.com/gin-gonic/gin"
)

// Admin CRUD Operations for Brands

// GetBrands - Get all brands (Admin)
func GetBrands(c *gin.Context) {
	var brands []models.Brand
	result := database.DB.Find(&brands)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brands"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    brands,
		"total":   len(brands),
	})
}

// GetBrand - Get brand by ID (Admin)
func GetBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand

	result := database.DB.First(&brand, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    brand,
	})
}

// CreateBrand - Create new brand (Admin)
func CreateBrand(c *gin.Context) {
	var req models.CreateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand := models.Brand{
		Name:        req.Name,
		Description: req.Description,
		Logo:        req.Logo,
	}

	result := database.DB.Create(&brand)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    brand,
		"message": "Brand created successfully",
	})
}

// UpdateBrand - Update brand (Admin)
func UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	var req models.UpdateBrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.First(&brand, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	// Update fields
	if req.Name != "" {
		brand.Name = req.Name
	}
	if req.Description != "" {
		brand.Description = req.Description
	}
	if req.Logo != "" {
		brand.Logo = req.Logo
	}
	if req.Status != "" {
		brand.Status = req.Status
	}

	database.DB.Save(&brand)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    brand,
		"message": "Brand updated successfully",
	})
}

// DeleteBrand - Delete brand (Admin)
func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand

	result := database.DB.First(&brand, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	database.DB.Delete(&brand)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Brand deleted successfully",
	})
}

// API for Flutter - Get Brand List
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

