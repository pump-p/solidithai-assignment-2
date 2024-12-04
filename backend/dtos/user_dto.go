package dtos

import "github.com/pump-p/solidithai-assignment-2/backend/models"

// UserResponse represents the user data exposed in API responses
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// MapToUserResponse converts a User model to a UserResponse DTO
func MapToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
