package handlers

import (
	"net/http"

	"gin-first/database"
	"gin-first/models"

	"github.com/gin-gonic/gin"
)

// Admin CRUD Operations for Stores

// GetStores - Get all stores (Admin)
func GetStores(c *gin.Context) {
	var stores []models.Store
	result := database.DB.Find(&stores)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stores"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stores,
		"total":   len(stores),
	})
}

// GetStore - Get store by ID (Admin)
func GetStore(c *gin.Context) {
	id := c.Param("id")
	var store models.Store

	result := database.DB.First(&store, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    store,
	})
}

// CreateStore - Create new store (Admin)
func CreateStore(c *gin.Context) {
	var req models.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := models.Store{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}

	result := database.DB.Create(&store)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    store,
		"message": "Store created successfully",
	})
}

// UpdateStore - Update store (Admin)
func UpdateStore(c *gin.Context) {
	id := c.Param("id")
	var store models.Store
	var req models.UpdateStoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.First(&store, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	// Update fields
	if req.Name != "" {
		store.Name = req.Name
	}
	if req.Address != "" {
		store.Address = req.Address
	}
	if req.Phone != "" {
		store.Phone = req.Phone
	}
	if req.Email != "" {
		store.Email = req.Email
	}
	if req.Status != "" {
		store.Status = req.Status
	}

	database.DB.Save(&store)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    store,
		"message": "Store updated successfully",
	})
}

// DeleteStore - Delete store (Admin)
func DeleteStore(c *gin.Context) {
	id := c.Param("id")
	var store models.Store

	result := database.DB.First(&store, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	database.DB.Delete(&store)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Store deleted successfully",
	})
}

