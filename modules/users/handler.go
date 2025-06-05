// modules/user/handler.go
package users

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// RegisterRoutes binds the user endpoints to Gin.
func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
    // Protect in prod
    rg.POST("/users", func(c *gin.Context) {
        var u User
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        created, err := svc.CreateUser(&u)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, created)
    })

    rg.GET("/users/me", func(c *gin.Context) {
        userID, exists := c.Get("userID") 
        if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

        user, err := svc.GetUserByID(userID.(string))	
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching sessions"})
            return
        }

        c.JSON(http.StatusOK, user)
    })

    rg.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        u, err := svc.GetUserByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, u)
    })
}
