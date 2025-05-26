package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/rlee603166/circl/modules/user"
)

func RegisterRoutes(rg *gin.RouterGroup, authSvc *Service, userSvc *user.Service) {
    rg.POST("/auth/google/log-in", func (c *gin.Context) {
        var body struct {
            Token string `json:"token"`
        }

        if err := c.BindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        authPayload, err := authSvc.VerifyGoogleToken(body.Token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        user, err := userSvc.GetUserByEmail(authPayload.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User service error"})
            return
        }

        token, err := authSvc.CreateToken(user.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "token": token,
            "user": user,
        })
    })
}
