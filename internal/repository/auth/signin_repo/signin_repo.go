package signin_repo

import (
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/gorm"
)

type SignInRepository struct {
	Database *gorm.DB
}

func NewSignInRepository(db *gorm.DB) *SignInRepository {
	return &SignInRepository{
		Database: db,
	}
}

func (repo *SignInRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
