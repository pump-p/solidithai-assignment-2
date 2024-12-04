package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pump-p/solidithai-assignment-2/backend/dtos"
	"github.com/pump-p/solidithai-assignment-2/backend/models"
	"github.com/pump-p/solidithai-assignment-2/backend/services"
)

// GetUsers retrieves all users
func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map all users to UserResponse DTOs
	var userResponses []dtos.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dtos.MapToUserResponse(&user))
	}

	c.JSON(http.StatusOK, userResponses)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Convert User model to UserResponse DTO
	userResponse := dtos.MapToUserResponse(user)

	c.JSON(http.StatusOK, userResponse)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert User model to UserResponse DTO
	userResponse := dtos.MapToUserResponse(&user)

	c.JSON(http.StatusCreated, gin.H{"user": userResponse, "message": "User registered successfully"})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.UpdateUser(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert User model to UserResponse DTO
	userResponse := dtos.MapToUserResponse(user)

	c.JSON(http.StatusOK, gin.H{"user": userResponse, "message": "User updated successfully"})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
