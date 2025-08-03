package auth_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/internal/router/auth/signin_router"
	"github.com/jokilagila/jokilagila-be/internal/router/auth/signup_router"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	signin_router.SetupSigninRouter(auth)
    signup_router.SetupSignUpRouter(auth)
}
