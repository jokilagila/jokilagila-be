package delete_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jokilagila/jokilagila-be/internal/service/user/delete_user_service"
	"github.com/jokilagila/jokilagila-be/pkg/response"
	"gorm.io/gorm"
)

type DeleteUserController struct {
    deleteUserService delete_user_service.DeleteUserService
}

func NewDeleteUserController(service delete_user_service.DeleteUserService) *DeleteUserController {
    return &DeleteUserController{
        deleteUserService: service,
    }
}

func (controller *DeleteUserController) DeleteUser(context *gin.Context) {
    id := context.Param("id")
    parsedID, err := uuid.Parse(id)
    if err != nil {
        context.JSON(http.StatusBadRequest, response.ErrorResponse{
            Success: false,
            Message: "Invalid user ID format",
        })
        return
    }

    err = controller.deleteUserService.DeleteUser(parsedID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            context.JSON(http.StatusNotFound, response.ErrorResponse{
                Success: false,
                Message: "User not found",
            })
        } else {
            context.JSON(http.StatusInternalServerError, response.ErrorResponse{
                Success: false,
                Message: "Failed to delete user",
            })
        }
        return
    }

    context.JSON(http.StatusOK, response.SuccessResponse{
        Success: true,
        Message: "User deleted successfully",
    })
}