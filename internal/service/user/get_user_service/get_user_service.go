package get_user_service

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/get_user_repo"
)

type GetUserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
}

type GetUserServiceImpl struct {
	GetUserRepository get_user_repo.GetUserRepository
}

func NewGetUserServiceImpl(repo get_user_repo.GetUserRepository) *GetUserServiceImpl {
	return &GetUserServiceImpl{
		GetUserRepository: repo,
	}
}

func (service *GetUserServiceImpl) GetAllUsers() ([]model.User, error) {
	return service.GetUserRepository.GetAllUsers()
}

func (service *GetUserServiceImpl) GetUserByID(id uuid.UUID) (*model.User, error) {
	return service.GetUserRepository.GetUserByID(id)
}
