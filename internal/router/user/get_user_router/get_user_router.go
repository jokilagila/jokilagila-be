package get_user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/controller/user/get_user_controller"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/get_user_repo"
	"github.com/jokilagila/jokilagila-be/internal/service/user/get_user_service"
	"github.com/jokilagila/jokilagila-be/middleware"
)

func SetupGetUserRouter(router *gin.RouterGroup) {
	db, _ := config.PostgresConfig()
	getUserRepo := get_user_repo.NewGetUserRepositoryImpl(db)
	getUserService := get_user_service.NewGetUserServiceImpl(getUserRepo)
	getUserController := get_user_controller.NewGetUserController(getUserService)

	router.GET("/users", middleware.JWTMiddleware(), middleware.AdminOnlyMiddleware(), getUserController.GetAllUsers)
	router.GET("/users/:id", middleware.JWTMiddleware(), middleware.AdminOnlyMiddleware(), getUserController.GetUserByID)
}