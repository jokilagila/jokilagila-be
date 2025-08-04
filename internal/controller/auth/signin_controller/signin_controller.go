package signin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/internal/dto"
	"github.com/jokilagila/jokilagila-be/internal/service/auth/signin_service"
	"github.com/jokilagila/jokilagila-be/pkg/response"
	"github.com/jokilagila/jokilagila-be/pkg/validator"
)

type SignInController struct {
	SignInService *signin_service.SignInService
}

func NewSignInController(service *signin_service.SignInService) *SignInController {
	return &SignInController{
		SignInService: service,
	}
}

func (controller *SignInController) SigninUser(context *gin.Context) {

	var request dto.UserLoginRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Errors:  validator.TranslateErrorMessage(err),
		})
		return
	}

	userResponse, err := controller.SignInService.Signin(request)

	if err != nil {
		context.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
			Errors:  map[string]string{"email": "Email atau password salah"},
		})
		return
	}

	context.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "User signed in successfully",
		Data:    userResponse,
	})

}
