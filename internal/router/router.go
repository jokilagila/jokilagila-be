package router

import (
	"log"

	"github.com/gin-gonic/gin"
	auth_routes "github.com/jokilagila/jokilagila-be/internal/router/auth"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	if err := route.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Gagal mengatur trusted proxies: %v", err)
	}

	api := route.Group("/api/v1")
	auth_routes.SetupAuthRoutes(api)

	return route
}
