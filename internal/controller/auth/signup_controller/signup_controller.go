package signup_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/internal/dto"
	"github.com/jokilagila/jokilagila-be/internal/service/auth/signup_service"
	"github.com/jokilagila/jokilagila-be/pkg/response"
	"github.com/jokilagila/jokilagila-be/pkg/validator"
)

type SignUpController struct {
	SignUpService *signup_service.SignUpService
}

func NewSignUpController(service *signup_service.SignUpService) *SignUpController {
	return &SignUpController{
		SignUpService: service,
	}
}

func (controller *SignUpController) SignupUser(context *gin.Context) {

	var request dto.UserCreateRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Errors:  validator.TranslateErrorMessage(err),
		})
		return
	}

	signUpResponse, err := controller.SignUpService.Signup(request)

	if err != nil {
		statusCode := http.StatusInternalServerError

		if err.Error() == "email sudah terdaftar" {
			statusCode = http.StatusConflict
		}

		context.JSON(statusCode, response.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data:    signUpResponse,
	})

}
