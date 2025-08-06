package edit_user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/controller/user/edit_user_controller"
	"github.com/jokilagila/jokilagila-be/internal/repository/user/edit_user_repo"
	"github.com/jokilagila/jokilagila-be/internal/service/user/edit_user_service"
	"github.com/jokilagila/jokilagila-be/middleware"
)

func SetupEditUserRouter(router *gin.RouterGroup) {
    db, _ := config.PostgresConfig()
    editRepo := edit_user_repo.NewEditUserRepositoryImpl(db)
    editService := edit_user_service.NewEditUserService(editRepo)
    editController := edit_user_controller.NewEditUserController(editService)

    router.PUT("/edit-user/:id", middleware.JWTMiddleware(), middleware.AdminOnlyMiddleware(), editController.EditUser)
}