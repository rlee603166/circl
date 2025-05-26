// modules/experience/handler.go
package experience

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
    rg.POST("/users/:id/experiences", func(c *gin.Context) {
        var e Experience
        if err := c.ShouldBindJSON(&e); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        e.UserID = c.Param("id")
        created, err := svc.CreateExperience(&e)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, created)
    })

    rg.GET("/users/:id/experiences", func(c *gin.Context) {
        list, err := svc.GetExperiencesByUserID(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, list)
    })

    rg.PUT("/experiences/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var e Experience
        if err := c.ShouldBindJSON(&e); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        e.ExperienceID = id
        updated, err := svc.UpdateExperience(&e)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, updated)
    })

    rg.DELETE("/experiences/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        if err := svc.DeleteExperience(id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    })
}
