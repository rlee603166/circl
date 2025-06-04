// main.go
package main

import (
    "log"
    "github.com/gin-gonic/gin"

    "github.com/rlee603166/circl/internal/db"
    "github.com/rlee603166/circl/internal/config"
    "github.com/rlee603166/circl/internal/middleware"

    "github.com/rlee603166/circl/modules/auth"
    "github.com/rlee603166/circl/modules/users"
    "github.com/rlee603166/circl/modules/messages"
    "github.com/rlee603166/circl/modules/sessions"
    "github.com/rlee603166/circl/modules/astralis"
    "github.com/rlee603166/circl/modules/educations"
    "github.com/rlee603166/circl/modules/experiences"

)

func main() {
    dbURL, port := config.Load()
    conn, err := db.Connect(dbURL)
    if err != nil {
        log.Fatalf("DB connection error: %v", err)
    }

    r := gin.Default()
    r.Use(middleware.CORS())

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "BREAK EVERYTHING!"})
    })


    // Middleware
    authSvc := auth.GetService()
    secured := r.Group("/api/v1", middleware.SecureHandler(authSvc))

    // User module
    uRepo   := users.NewRepository(conn)
    uSvc    := users.NewService(uRepo)

    // Experience module
    expRepo := experiences.NewRepository(conn)
    expSvc  := experiences.NewService(expRepo)

    // Education module
    edRepo  := educations.NewRepository(conn)
    edSvc   := educations.NewService(edRepo)

    // Session module
    seRepo  := sessions.NewRepository(conn)
    seSvc   := sessions.NewService(seRepo)

    // Message module
    mRepo   := messages.NewRepository(conn)
    mSvc    := messages.NewService(mRepo)

    // Route registration
    auth.RegisterRoutes(r, authSvc, uSvc)
    users.RegisterRoutes(secured, uSvc)
    experiences.RegisterRoutes(secured, expSvc)
    educations.RegisterRoutes(secured, edSvc)
    sessions.RegisterRoutes(secured, seSvc, mSvc)
    messages.RegisterRoutes(secured, mSvc)
    astralis.RegisterRoutes(secured)

    r.Run(port)
}
