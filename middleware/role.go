package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func AdminOnlyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        roleVal, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: No role found"})
            c.Abort()
            return
        }

        role, ok := roleVal.(string)
        if !ok || strings.ToLower(role) != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin access only"})
            c.Abort()
            return
        }

        c.Next()
    }
}
