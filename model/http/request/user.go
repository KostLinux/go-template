package request

// CreateUserRequest represents the create user request structure
type CreateUser struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"secretpassword123"`
}

// UpdateUserRequest represents the update user request structure
type UpdateUser struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password,omitempty" example:"newpassword123"`
}
