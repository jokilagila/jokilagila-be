package router

import (
	"github.com/gin-gonic/gin"
	auth_routes "github.com/jokilagila/jokilagila-be/internal/router/auth"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	api := route.Group("/api/v1")
	auth_routes.SetupAuthRoutes(api)

	return route
}
