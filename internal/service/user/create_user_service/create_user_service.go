package create_user_service

import (
	"errors"

	"github.com/jokilagila/jokilagila-be/internal/dto"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/create_user_repo"
	"github.com/jokilagila/jokilagila-be/pkg/hashpassword"
)

type CreateUserService struct {
	CreateUserRepository *create_user_repo.CreateUserRepository
}

func NewCreateUserService(repo *create_user_repo.CreateUserRepository) *CreateUserService {
	return &CreateUserService{
		CreateUserRepository: repo,
	}
}

func (service *CreateUserService) CreateUser(request dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	if service.CreateUserRepository.IsEmailExists(request.Email) {
        return nil, errors.New("email sudah terdaftar")
    }

    if request.Password != request.ConfirmPassword {
        return nil, errors.New("password dan konfirmasi password tidak cocok")
    }

    hashedPassword, err := hashpassword.HashPassword(request.Password)
    if err != nil {
        return nil, err
    }

    user := model.User{
        Name:     request.Name,
        Email:    request.Email,
        Password: hashedPassword,
        Role:     "user",
    }

    if err := service.CreateUserRepository.CreateUser(&user); err != nil {
        return nil, err
    }

    response := &dto.UserCreateResponse{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        Role:      user.Role,
        Phone:     user.Phone,
        CreatedAt: user.CreatedAt.String(),
        UpdatedAt: user.UpdatedAt.String(),
    }

    return response, nil
}
