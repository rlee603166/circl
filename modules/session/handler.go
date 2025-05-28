package session

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
    rg.GET("/session", func (c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
            return
        }

        sessions, err := svc.GetSessionsByUserID(userID.(string))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching sessions"})
            return
        }

        c.JSON(http.StatusOK, sessions)
    })

    rg.POST("/session", func (c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
            return
        }       
        sessionID := uuid.New().String()
        newSession := &CreateSession{
            UserID: userID.(string),
            SessionID: sessionID,
        }
        
        created,err := svc.CreateSession(newSession)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating session"})
            return
        }

        c.JSON(http.StatusCreated, created)
    })

    rg.GET("/session/:id", func (c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
            return
        }

        id := c.Param("id")
        s, err := svc.GetSessionByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        if s.UserID != userID.(string) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to session"})
            return
        }

        c.JSON(http.StatusOK, s)
    })
}
