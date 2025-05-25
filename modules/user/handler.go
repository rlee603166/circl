// modules/user/handler.go
package user

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// RegisterRoutes binds the user endpoints to Gin.
func RegisterRoutes(r *gin.RouterGroup, svc *Service) {
    r.POST("/users", func(c *gin.Context) {
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

    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        u, err := svc.GetUserByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, u)
    })
}
