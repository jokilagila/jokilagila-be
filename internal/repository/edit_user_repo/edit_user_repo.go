package edit_user_repo

import (
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type EditUserRepository struct {
	Database *gorm.DB
}

func NewEditUserRepository(db *gorm.DB) *EditUserRepository {
	return &EditUserRepository{
		Database: db,
	}
}

func (repo *EditUserRepository) UpdateUser(id uint, userData map[string]interface{}) error {
	var user model.User
	err := repo.Database.First(&user, id).Error
	if err != nil {
		return err
	}

	err = repo.Database.Model(&user).Updates(userData).Error
	if err != nil {
		return err
	}

	return nil
}
