package signin_router

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/controller/signin_controller"
	"github.com/jokilagila/jokilagila-be/internal/repository/signin_repo"
	"github.com/jokilagila/jokilagila-be/internal/service/signin_service"
	"github.com/jokilagila/jokilagila-be/middleware"
)

func SetupSigninRouter(router *gin.RouterGroup) {
	db, _ := config.PostgresConfig()
	signinRepo := signin_repo.NewSignInRepository(db)
	signinService := signin_service.NewSignInService(signinRepo)
	signinController := signin_controller.NewSignInController(signinService)

	router.POST("/signin", middleware.RateLimiterMiddleware(), signinController.SigninUser)
}
