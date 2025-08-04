package get_user_repo

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type GetUserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
}

type GetUserRepositoryImpl struct {
	Database *gorm.DB
}

func NewGetUserRepositoryImpl(db *gorm.DB) *GetUserRepositoryImpl {
	return &GetUserRepositoryImpl{
		Database: db,
	}
}

func (repo *GetUserRepositoryImpl) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := repo.Database.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *GetUserRepositoryImpl) GetUserByID(id uint) (*model.User, error) {
	var user model.User

	if err := repo.Database.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
