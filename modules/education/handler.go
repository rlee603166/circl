// modules/education/handler.go
package education

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, svc *Service) {
    rg.POST("/users/:id/educations", func(c *gin.Context) {
        var e CreateEducation 
        if err := c.ShouldBindJSON(&e); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        e.UserID = c.Param("id")
        created, err := svc.CreateEducation(&e)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, created)
    })

    rg.GET("/users/:id/educations", func(c *gin.Context) {
        list, err := svc.GetEducationsByUserID(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, list)
    })

    rg.PUT("/educations/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var e Education
        if err := c.ShouldBindJSON(&e); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        e.EducationID = id
        updated, err := svc.UpdateEducation(&e)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, updated)
    })

    rg.DELETE("/educations/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        if err := svc.DeleteEducation(id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    })
}
