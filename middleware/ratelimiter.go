package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(5.0/300.0, 5)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !limiter.Allow() {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Terlalu banyak permintaan, silakan coba lagi dalam beberapa menit",
			})
			return
		}
		context.Next()
	}
}
