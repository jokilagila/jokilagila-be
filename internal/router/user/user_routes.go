package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/internal/router/user/create_user_router"
)

func SetupUserRoutes(router *gin.RouterGroup) {
    user := router.Group("/user")
    create_user_router.SetupCreateUserRouter(user)
}