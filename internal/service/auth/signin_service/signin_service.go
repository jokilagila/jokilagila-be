package signin_service

import (
	"errors"

	"github.com/jokilagila/jokilagila-be/internal/dto"
	"github.com/jokilagila/jokilagila-be/internal/repository/auth/signin_repo"
	"github.com/jokilagila/jokilagila-be/pkg/generatejwt"
	"golang.org/x/crypto/bcrypt"
)

type SignInService struct {
	SignInRepository *signin_repo.SignInRepository
}

func NewSignInService(repo *signin_repo.SignInRepository) *SignInService {
	return &SignInService{
		SignInRepository: repo,
	}
}

func (service *SignInService) Signin(request dto.UserLoginRequest) (*dto.UserResponse, error) {
	user, err := service.SignInRepository.FindUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("password salah")
	}

	jsonWebToken, err := generatejwt.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return nil, errors.New("gagal membuat token")
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Token:     jsonWebToken,
	}, nil
}
