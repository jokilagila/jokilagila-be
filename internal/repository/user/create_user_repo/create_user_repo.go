package create_user_repo

import (
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type CreateUserRepository struct {
    Database *gorm.DB
}

func NewCreateUserRepository(db *gorm.DB) *CreateUserRepository {
    return &CreateUserRepository{
        Database: db,
    }
}

func (repo *CreateUserRepository) CreateUser(user *model.User) error {
    return repo.Database.Create(user).Error
}

func (repo *CreateUserRepository) IsEmailExists(email string) bool {
    var user model.User
    err := repo.Database.Where("email = ?", email).First(&user).Error

    if err == nil {
        return true
    }

    if err == gorm.ErrRecordNotFound {
        return false
    }

    return false
}