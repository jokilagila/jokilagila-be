package create_user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/controller/user/create_user_controller"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/create_user_repo"
	"github.com/jokilagila/jokilagila-be/internal/service/user/create_user_service"
	"github.com/jokilagila/jokilagila-be/middleware"
)

func SetupCreateUserRouter(router *gin.RouterGroup) {
	db, _ := config.PostgresConfig()
	createUserRepo := create_user_repo.NewCreateUserRepository(db)
	createUserService := create_user_service.NewCreateUserService(createUserRepo)
	createUserController := create_user_controller.NewCreateUserController(*createUserService)

	router.POST("/create-user", middleware.JWTMiddleware(), middleware.AdminOnlyMiddleware(), createUserController.CreateUser)
}
