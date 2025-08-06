package delete_user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/controller/user/delete_user_controller"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/delete_user_repo"
	"github.com/jokilagila/jokilagila-be/internal/service/user/delete_user_service"
	"github.com/jokilagila/jokilagila-be/middleware"
)

func SetupDeleteUserRouter(router *gin.RouterGroup) {
	db, _ := config.PostgresConfig()
	deleteRepo := delete_user_repo.NewDeleteUserRepositoryImpl(db)
	deleteService := delete_user_service.NewDeleteUserService(deleteRepo)
	deleteController := delete_user_controller.NewDeleteUserController(deleteService)

	router.DELETE("/delete-user/:id", middleware.JWTMiddleware(), middleware.AdminOnlyMiddleware(), deleteController.DeleteUser)
}