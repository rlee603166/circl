package sessions

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rlee603166/circl/modules/messages"
)

func RegisterRoutes(rg *gin.RouterGroup, sessionSvc *Service, msgSvc *messages.Service) {
	rg.GET("/sessions", func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		sessions, err := sessionSvc.GetSessionsByUserID(userID.(string))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching sessions"})
			return
		}

		c.JSON(http.StatusOK, sessions)
	})

	rg.POST("/sessions", func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

        newChat := ""
		sessionID := uuid.New().String()
		newSession := &CreateSession{
			UserID:    userID.(string),
			SessionID: sessionID,
			Title: &newChat,
		}

		created, err := sessionSvc.CreateSession(newSession)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating session"})
			return
		}

		c.JSON(http.StatusCreated, created)
	})

	rg.GET("/sessions/:id", func(c *gin.Context) {
		// userID, exists := c.Get("userID")
		// if !exists {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		// 	return
		// }

		id := c.Param("id")
		s, err := sessionSvc.GetSessionByID(id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}

		// if s.UserID != userID.(string) {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to session"})
		// 	return
		// }

		c.JSON(http.StatusOK, s)
	})

	rg.GET("/sessions/:id/messages", func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		id := c.Param("id")
		fmt.Println(id)
		session, err := sessionSvc.GetSessionByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}

		if session.UserID != userID.(string) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
			return
		}

		messages, err := msgSvc.GetMessagesBySessionID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Messages not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"session":  session,
			"messages": messages,
		})
	})
}
