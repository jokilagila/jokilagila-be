package edit_user_service

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/edit_user_repo"
)

type EditUserService interface {
	UpdateUser(id uuid.UUID, userData map[string]interface{}) error
}

type editUserService struct {
	repo edit_user_repo.EditUserRepository
}

var _ EditUserService = &editUserService{}

func NewEditUserService(r edit_user_repo.EditUserRepository) EditUserService {
	return &editUserService{repo: r}
}

func (s *editUserService) UpdateUser(id uuid.UUID, userData map[string]interface{}) error {
	return s.repo.UpdateUser(id, userData)
}
