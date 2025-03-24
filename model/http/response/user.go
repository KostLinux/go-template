package response

import (
	"time"
)

// UserResponse represents the user response structure
type User struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-03-24T15:04:05Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-03-24T15:04:05Z"`
}

// UserListResponse represents a list of users
type UserList struct {
	Users []User `json:"users"`
}

type DeleteUser struct {
	Message string `json:"message" example:"user deleted successfully"`
}

type UserErrorBadRequest struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"invalid id for user"`
	Error   string `json:"error,omitempty" example:"User ID must be a positive integer"`
}

type UserErrorNotFound struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"user not found"`
	Error   string `json:"error,omitempty" example:"User with ID 1 not found"`
}

type UserErrorConflict struct {
	Code    int    `json:"code" example:"409"`
	Message string `json:"message" example:"user already exists"`
	Error   string `json:"error,omitempty" example:"User with email already exists"`
}

type UserErrorInternalServer struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"failed to create user"`
	Error   string `json:"error,omitempty" example:"Internal server error"`
}
