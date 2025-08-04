package edit_user_repo

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type EditUserRepository interface {
	UpdateUser(id uuid.UUID, userData map[string]interface{}) error
}

type EditUserRepositoryImpl struct {
	db *gorm.DB
}

var _ EditUserRepository = &EditUserRepositoryImpl{}

func NewEditUserRepositoryImpl(db *gorm.DB) *EditUserRepositoryImpl {
	return &EditUserRepositoryImpl{db: db}
}

func (repo *EditUserRepositoryImpl) UpdateUser(id uuid.UUID, userData map[string]interface{}) error {
	result := repo.db.Model(&model.User{}).Where("id = ?", id).Updates(userData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
