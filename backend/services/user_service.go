package services

import (
	"errors"

	"github.com/pump-p/solidithai-assignment-2/backend/database"
	"github.com/pump-p/solidithai-assignment-2/backend/models"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	return users, nil
}

// GetUserByID retrieves a specific user by ID
func GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// CreateUser creates a new user
func CreateUser(user *models.User) error {
	// Hash the user's password
	if err := user.HashPassword(); err != nil {
		return errors.New("failed to hash password")
	}

	// Save the user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

// UpdateUser updates an existing user
func UpdateUser(id string, updates map[string]interface{}) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Apply updates
	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return &user, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id string) error {
	if err := database.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
