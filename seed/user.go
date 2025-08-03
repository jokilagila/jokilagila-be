package seed

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"github.com/jokilagila/jokilagila-be/pkg/hashpassword"
	"gorm.io/gorm"
)

func UserSeed() error {
	postgresDB, err := config.PostgresConfig()
	if err != nil {
		return err
	}

	var existingAdmin model.User
	err = postgresDB.Where("role = ?", "admin").First(&existingAdmin).Error
    
	if err == nil {
		return nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := hashpassword.HashPassword("minjokkeren123")
	if err != nil {
		return err
	}

	phoneNumber := "081217952403"
	adminData := model.User{
		ID:       uuid.New(),
		Name:     "Admin",
		Email:    "minjokkeren@mail.com",
		Phone:    &phoneNumber,
		Role:     "admin",
		Password: hash,
	}

	if err := postgresDB.Create(&adminData).Error; err != nil {
		return err
	}

	return nil
}
