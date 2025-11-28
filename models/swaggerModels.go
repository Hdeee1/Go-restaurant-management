package models

import "time"

// UserResponse represents the user data returned in API responses
type UserResponse struct {
	ID           uint      `json:"id" example:"1"`
	CreatedAt    time.Time `json:"created_at" example:"2025-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2025-01-01T00:00:00Z"`
	Role         string    `json:"role" example:"admin"`
	FirstName    string    `json:"first_name" example:"John"`
	LastName     string    `json:"last_name" example:"Doe"`
	Email        string    `json:"email" example:"user@example.com"`
	Avatar       string    `json:"avatar" example:"https://example.com/avatar.jpg"`
	Phone        string    `json:"phone" example:"+1234567890"`
	Token        string    `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string    `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	UserID       string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// SignUpRequest represents the request body for user registration
type SignUpRequest struct {
	Email     string `json:"email" example:"user@example.com" binding:"required,email"`
	Password  string `json:"password" example:"Passw0rd!" binding:"required,min=8"`
	FirstName string `json:"first_name" example:"John" binding:"required,min=2,max=100"`
	LastName  string `json:"last_name" example:"Doe" binding:"required,min=2,max=100"`
	Phone     string `json:"phone" example:"+1234567890" binding:"required"`
	Avatar    string `json:"avatar" example:"https://example.com/avatar.jpg"`
	Role      string `json:"role" example:"user|admin" binding:"required"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com" binding:"required,email"`
	Password string `json:"password" example:"Passw0rd!" binding:"required"`
}

// LoginResponse represents the response body after successful login
type LoginResponse struct {
	Message string `json:"message" example:"Login successfully"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	UserID  string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// SignUpResponse represents the response body after successful registration
type SignUpResponse struct {
	Message string `json:"message" example:"Successfully signed up"`
	UserID  string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

// UsersListResponse represents paginated users list
type UsersListResponse struct {
	Data  []UserResponse `json:"data"`
	Page  string         `json:"page" example:"1"`
	Limit string         `json:"limit" example:"10"`
}
