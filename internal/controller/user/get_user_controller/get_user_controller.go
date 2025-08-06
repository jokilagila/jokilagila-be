package get_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/service/user/get_user_service"
	"github.com/jokilagila/jokilagila-be/pkg/response"
)

type GetUserController struct {
	GetUserService get_user_service.GetUserService
}

func NewGetUserController(service get_user_service.GetUserService) *GetUserController {
	return &GetUserController{
		GetUserService: service,
	}
}

func (controller *GetUserController) GetAllUsers(context *gin.Context) {
	users, err := controller.GetUserService.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve users",
		})
	} else {
		context.JSON(http.StatusOK, response.SuccessResponse{
			Success: true,
			Message: "Users retrieved successfully",
			Data:    users,
		})
	}
}
func (controller *GetUserController) GetUserByID(context *gin.Context) {
	userID := context.Param("id")
	uuidID, err := uuid.Parse(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Invalid user ID format",
		})
		return
	}
	user, err := controller.GetUserService.GetUserByID(uuidID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve user",
		})
	} else if user == nil {
		context.JSON(http.StatusNotFound, response.ErrorResponse{
			Success: false,
			Message: "User not found",
		})
	} else {
		context.JSON(http.StatusOK, response.SuccessResponse{
			Success: true,
			Message: "User retrieved successfully",
			Data:    user,
		})
	}
}
