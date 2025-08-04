package delete_user_repo

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type DeleteUserRepository interface {
	DeleteUser(id uuid.UUID) error
}

type DeleteUserRepositoryImpl struct {
	db *gorm.DB
}

var _ DeleteUserRepository = &DeleteUserRepositoryImpl{}

func NewDeleteUserRepositoryImpl(db *gorm.DB) *DeleteUserRepositoryImpl {
	return &DeleteUserRepositoryImpl{db: db}
}

func (repo *DeleteUserRepositoryImpl) DeleteUser(id uuid.UUID) error {
	result := repo.db.Delete(&model.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
