package signup_router

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/controller/signup_controller"
	"github.com/jokilagila/jokilagila-be/internal/repository/signup_repo"
	"github.com/jokilagila/jokilagila-be/internal/service/signup_service"
)

func SetupSignUpRouter(router *gin.RouterGroup) {
	db, _ := config.PostgresConfig()
	signUpRepo := signup_repo.NewSignUpRepository(db)
	signUpService := signup_service.NewSignUpService(signUpRepo)
	signUpController := signup_controller.NewSignUpController(signUpService)

	router.POST("/signup", signUpController.SignupUser)
}
