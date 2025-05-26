package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/rlee603166/circl/modules/auth"
)


func SecureHandler(s *auth.Service) gin.HandlerFunc {

    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
            return
        }
        
        idToken := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := s.VerifyToken(idToken)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
            return
        }

        username, ok := claims["username"].(string)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
            return
        }

        c.Set("username", username)
        c.Next()
    }
}
