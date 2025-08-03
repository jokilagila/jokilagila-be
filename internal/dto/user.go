package dto

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Phone     *string   `json:"phone,omitempty"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Token     string    `json:"token"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

type UserCreateResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
