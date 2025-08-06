package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jokilagila/jokilagila-be/internal/router/user/create_user_router"
	"github.com/jokilagila/jokilagila-be/internal/router/user/edit_user_router"
	"github.com/jokilagila/jokilagila-be/internal/router/user/get_user_router"
)

func SetupUserRoutes(router *gin.RouterGroup) {
    user := router.Group("/user")
    create_user_router.SetupCreateUserRouter(user)
    get_user_router.SetupGetUserRouter(user)
    edit_user_router.SetupEditUserRouter(user)
}