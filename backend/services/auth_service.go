package services

import (
	"errors"

	"github.com/pump-p/solidithai-assignment-2/backend/database"
	"github.com/pump-p/solidithai-assignment-2/backend/models"
	"github.com/pump-p/solidithai-assignment-2/backend/utils"
	"gorm.io/gorm"
)

// Signup registers a new user
func Signup(user *models.User) (string, error) {
	// Check if the email already exists
	var existingUser models.User
	err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return "", errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("failed to check existing email")
	}

	// Hash the password
	if err := user.HashPassword(); err != nil {
		return "", errors.New("failed to hash password")
	}

	// Save the user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return "", errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// Login authenticates a user and generates a JWT token
func Login(email, password string) (string, error) {
	// Retrieve the user by email
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", errors.New("failed to retrieve user")
	}

	// Check if the provided password matches
	if !user.CheckPassword(password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
