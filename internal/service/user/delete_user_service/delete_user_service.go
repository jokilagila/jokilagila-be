package delete_user_service

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/delete_user_repo"
)

type DeleteUserService interface {
	DeleteUser(id uuid.UUID) error
}

type deleteUserService struct {
	repo delete_user_repo.DeleteUserRepository
}

var _ DeleteUserService = &deleteUserService{}

func NewDeleteUserService(r delete_user_repo.DeleteUserRepository) DeleteUserService {
	return &deleteUserService{repo: r}
}

func (s *deleteUserService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
