package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Password string    `json:"password"`
	Role     string    `json:"role" gorm:"default:user"`
	Phone    *string   `json:"phone"`
}
