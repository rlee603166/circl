package auth

import (
    "os"
    "context"
    "net/http"
    "strings"

    "google.golang.org/api/idtoken"
    "github.com/gin-gonic/gin"
)

var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")

func GoogleAuthMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
            return
        }

        idToken := strings.TrimPrefix(authHeader, "Bearer ")

        payload, err := idtoken.Validate(context.Background(), idToken, googleClientID)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Set("userEmail", payload.Claims["email"])
        c.Set("userName", payload.Claims["name"])
        c.Next()
    }
}
