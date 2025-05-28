// main.go
package main

import (
    "log"
    "github.com/gin-gonic/gin"

    "github.com/rlee603166/circl/internal/db"
    "github.com/rlee603166/circl/internal/config"
    "github.com/rlee603166/circl/internal/middleware"

    "github.com/rlee603166/circl/modules/auth"
    "github.com/rlee603166/circl/modules/user"
    "github.com/rlee603166/circl/modules/message"
    "github.com/rlee603166/circl/modules/session"
    "github.com/rlee603166/circl/modules/education"
    "github.com/rlee603166/circl/modules/experience"

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
    uRepo   := user.NewRepository(conn)
    uSvc    := user.NewService(uRepo)

    // Experience module
    expRepo := experience.NewRepository(conn)
    expSvc  := experience.NewService(expRepo)

    // Education module
    edRepo  := education.NewRepository(conn)
    edSvc   := education.NewService(edRepo)

    // Session module
    seRepo  := session.NewRepository(conn)
    seSvc   := session.NewService(seRepo)

    // Message module
    mRepo   := message.NewRepository(conn)
    mSvc    := message.NewService(mRepo)

    // Route registration
    auth.RegisterRoutes(r, authSvc, uSvc)
    user.RegisterRoutes(secured, uSvc)
    experience.RegisterRoutes(secured, expSvc)
    education.RegisterRoutes(secured, edSvc)
    session.RegisterRoutes(secured, seSvc)
    message.RegisterRoutes(secured, mSvc)

    r.Run(port)
}
