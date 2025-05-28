package message

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
    rg.GET("/message", func (c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
            return
        }

        messages, err := svc.GetMessagesByUserID(userID.(string))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching sessions"})
            return
        }

        c.JSON(http.StatusOK, messages)
    })

    // rg.GET("/message/:id", func (c *gin.Context) {
    //     userID, exists := c.Get("userID")
    //     if !exists {
    //         c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
    //         return
    //     }
    //
    //     id := c.Param("id")
    //     msg, err := svc.GetMessageByID(id)
    //     if err != nil {
    //         c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    //         return
    //     }
    //
    //     if msg.UserID != userID.(string) {
    //         c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to session"})
    //         return
    //     }
    //
    //     c.JSON(http.StatusOK, msg)
    // })
}
