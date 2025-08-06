package edit_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/service/user/edit_user_service"
	"github.com/jokilagila/jokilagila-be/pkg/response"
)

type EditUserController struct {
	editUserService edit_user_service.EditUserService
}

func NewEditUserController(service edit_user_service.EditUserService) *EditUserController {
	return &EditUserController{
		editUserService: service,
	}
}

func (controller *EditUserController) EditUser(context *gin.Context) {
	id := context.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Invalid user ID format",
		})
		return
	}

	var userData map[string]interface{}
	if context.Request.ContentLength > 0 {
		if err := context.ShouldBindJSON(&userData); err != nil {
			context.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success: false,
				Message: "Invalid request data",
			})
			return
		}
	}

	updatedUser, err := controller.editUserService.UpdateUser(parsedID, userData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	context.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    updatedUser,
	})

}
