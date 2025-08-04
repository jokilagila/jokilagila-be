package get_user_repo

import (
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type GetUserRepository struct {
	Database *gorm.DB
}

func NewGetUserRepository(db *gorm.DB) *GetUserRepository {
	return &GetUserRepository{
		Database: db,
	}
}

func (repo *GetUserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := repo.Database.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *GetUserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := repo.Database.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
