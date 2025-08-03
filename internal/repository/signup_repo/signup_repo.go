package signup_repo

import (
    "github.com/jokilagila/jokilagila-be/internal/model"
    "gorm.io/gorm"
)

type SignUpRepository struct {
    Database *gorm.DB
}

func NewSignUpRepository(db *gorm.DB) *SignUpRepository {
    return &SignUpRepository{
        Database: db,
    }
}

func (repo *SignUpRepository) SignupUser(user *model.User) error {
    return repo.Database.Create(user).Error
}

func (repo *SignUpRepository) IsEmailExists(email string) bool {
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
