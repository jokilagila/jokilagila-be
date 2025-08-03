package signup_service

import (
	"errors"

	"github.com/jokilagila/jokilagila-be/internal/dto"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"github.com/jokilagila/jokilagila-be/internal/repository/signup_repo"
	"github.com/jokilagila/jokilagila-be/pkg/hashpassword"
)

type SignUpService struct {
	SignUpRepository *signup_repo.SignUpRepository
}

func NewSignUpService(repo *signup_repo.SignUpRepository) *SignUpService {
	return &SignUpService{
		SignUpRepository: repo,
	}
}

func (service *SignUpService) Signup(request dto.UserCreateRequest) (*dto.UserCreateResponse, error) {

	if service.SignUpRepository.IsEmailExists(request.Email) {
		return nil, errors.New("email sudah terdaftar")
	}

	if request.Password != request.ConfirmPassword {
		return nil, errors.New("password dan konfirmasi password tidak cocok")
	}

	hashedPassword, err := hashpassword.HashPassword(request.Password)

	if err != nil {
		return nil, err
	}

	if request.Password != request.ConfirmPassword {
		return nil, errors.New("password dan konfirmasi password tidak cocok")
	}

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := service.SignUpRepository.SignupUser(&user); err != nil {
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
