package delete_user_repo

import (
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type DeleteUserRepository struct {
	Database *gorm.DB
}

func NewDeleteUserRepository(db *gorm.DB) *DeleteUserRepository {
	return &DeleteUserRepository{
		Database: db,
	}
}

func (repo *DeleteUserRepository) DeleteUserByID(id uint) error {
	return repo.Database.Delete(&model.User{}, id).Error
}
