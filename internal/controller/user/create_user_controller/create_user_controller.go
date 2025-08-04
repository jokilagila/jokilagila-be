package create_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/internal/dto"
	"github.com/jokilagila/jokilagila-be/internal/service/user/create_user_service"
	"github.com/jokilagila/jokilagila-be/pkg/response"
	"github.com/jokilagila/jokilagila-be/pkg/validator"
)

type CreateUserController struct {
	CreateUserService create_user_service.CreateUserService
}

func NewCreateUserController(service create_user_service.CreateUserService) *CreateUserController {
	return &CreateUserController{
		CreateUserService: service,
	}
}

func (controller *CreateUserController) CreateUser(context *gin.Context) {
	var req dto.UserCreateRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Errors:  validator.TranslateErrorMessage(err),
		})
		return
	}

	createUserResponse, err := controller.CreateUserService.CreateUser(req)

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
        Data:    createUserResponse,
    });
}
