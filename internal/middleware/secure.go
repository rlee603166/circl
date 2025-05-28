package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/rlee603166/circl/modules/auth"
)


func SecureHandler(authSvc *auth.Service) gin.HandlerFunc {

    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
            return
        }
        
        accessToken := strings.TrimPrefix(authHeader, "Bearer ")
        tokenPayload, err := authSvc.VerifyAccessToken(accessToken)
        if err != nil  {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
            return
        }

        c.Set("userID", *tokenPayload.UserID)
        c.Set("email", *tokenPayload.Email)
        c.Next()
    }
}
