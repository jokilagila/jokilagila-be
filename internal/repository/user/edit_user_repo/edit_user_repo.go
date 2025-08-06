package edit_user_repo

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type EditUserRepository interface {
	UpdateUser(id uuid.UUID, userData map[string]interface{}) (*model.User, error)
}

type EditUserRepositoryImpl struct {
	db *gorm.DB
}

func NewEditUserRepositoryImpl(db *gorm.DB) *EditUserRepositoryImpl {
	return &EditUserRepositoryImpl{db: db}
}

func (repo *EditUserRepositoryImpl) UpdateUser(id uuid.UUID, userData map[string]interface{}) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := repo.db.Model(&user).Updates(userData).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
