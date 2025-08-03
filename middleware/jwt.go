package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JWTMiddleware() gin.HandlerFunc {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not found in .env file")
	}

	return func(context *gin.Context) {
		stringToken := context.GetHeader("Authorization")

		if stringToken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}

		stringToken = strings.TrimPrefix(stringToken, "Bearer ")

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(stringToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{
					Err:  http.ErrNotSupported,
					Type: gin.ErrorTypePublic,
				}
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			context.Abort()
			return
		}

		context.Set("email", claims.Subject)

		context.Next()
	}
}
