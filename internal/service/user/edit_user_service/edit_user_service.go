package edit_user_service

import (
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/edit_user_repo"
)

type EditUserService interface {
	UpdateUser(id uuid.UUID, userData map[string]interface{}) (*model.User, error)
}

type editUserService struct {
	repo edit_user_repo.EditUserRepository
}

func NewEditUserService(repo edit_user_repo.EditUserRepository) EditUserService {
    return &editUserService{
        repo: repo,
    }
}

func (s *editUserService) UpdateUser(id uuid.UUID, userData map[string]interface{}) (*model.User, error) {
	return s.repo.UpdateUser(id, userData)
}
