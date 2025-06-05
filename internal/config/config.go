// internal/config/config.go
package config

import (
    // "fmt"
    "log"
    "os"

    _ "github.com/joho/godotenv/autoload"
)

// Load reads env vars, constructs a PostgreSQL URL, and returns config values.
func Load() (dbConn string, serverPort string) {
    // user := os.Getenv("DB_USER")
    // // pass := os.Getenv("DB_PASSWORD")
    // host := os.Getenv("DB_HOST")
    // port := os.Getenv("DB_PORT")
    // name := os.Getenv("DB_NAME")
    //
    // if user == "" || host == "" || port == "" || name == "" {
    //     log.Fatal("Missing one or more DB env vars: DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME")
    // }

    url := os.Getenv("DB_URL")

    if url == "" {
        log.Fatal("Missing DB_URL")
    }

    serverPort = os.Getenv("PORT")
    if serverPort == "" {
        serverPort = ":8080"
    }

    return url, serverPort
}

