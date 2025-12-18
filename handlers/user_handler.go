package handlers

import (
	"net/http"
	"strconv"

	"gin-first/database"
	"gin-first/models"

	"github.com/gin-gonic/gin"
)

// Admin CRUD Operations for Users

// GetUsers - Get all users (Admin)
func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
		"total":   len(users),
	})
}

// GetUser - Get user by ID (Admin)
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

// CreateUser - Create new user (Admin)
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    user,
		"message": "User created successfully",
	})
}

// UpdateUser - Update user (Admin)
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var req models.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Status != "" {
		user.Status = req.Status
	}

	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
		"message": "User updated successfully",
	})
}

// DeleteUser - Delete user (Admin)
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

// API for Flutter - Get User List
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

