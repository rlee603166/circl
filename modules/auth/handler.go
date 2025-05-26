package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/rlee603166/circl/modules/user"
)

func RegisterRoutes(r *gin.Engine, authSvc *Service, userSvc *user.Service) {
    r.POST("/auth/log-in", func (c *gin.Context) {
        var body struct {
            Email       string `json:"email"`
            Password    string `json:"password"`
        }

        if err := c.BindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        user, err := userSvc.GetUserByEmail(body.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User service error"})
            return
        }

        // if user.HashedPassword == nil || bcrypt.CompareHashAndPassword([]byte(*user.HashedPassword), []byte(body.Password)) != nil {
        //     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        //     return
        // }

        if user.HashedPassword == nil || *user.HashedPassword != body.Password {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
            return
        }

        accessToken, err := authSvc.CreateAccessToken(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Access Token generation error"})
            return
        }

        refreshToken, err := authSvc.CreateRefreshToken(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Refresh Token generation error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "accessToken": accessToken,
            "refreshToken": refreshToken,
            "user": gin.H{
                "userID": user.UserID,
                "firstName": user.FirstName, 
                "lastName": user.LastName,
                "email": user.Email,
                "pfpURL": user.PfpURL,
            },
        })
    })

    r.POST("/auth/google/log-in", func (c *gin.Context) {
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

        accessToken, err := authSvc.CreateAccessToken(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Access Token generation error"})
            return
        }

        refreshToken, err := authSvc.CreateRefreshToken(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Refresh Token generation error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "accessToken": accessToken,
            "refreshToken": refreshToken,
            "user": gin.H{
                "userID": user.UserID,
                "firstName": user.FirstName, 
                "lastName": user.LastName,
                "email": user.Email,
                "pfpURL": user.PfpURL,
            },
        })
    })

    r.POST("/auth/refresh", func (c *gin.Context) {
        var body struct {
            refreshToken    string `json:"refreshToken"`
        }

        if err := c.BindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }
        
        tokenPayload, err := authSvc.VerifyRefreshToken(body.refreshToken)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
            return
        }

        user, err := userSvc.GetUserByEmail(tokenPayload.Email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User service error"})
            return
        }

        accessToken, err := authSvc.CreateAccessToken(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Access Token generation error"})
            return
        }

        refreshToken, err := authSvc.CreateRefreshToken(user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Refresh Token generation error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "accessToken": accessToken,
            "refreshToken": refreshToken,
        })    
    })
}
