// internal/db/postgres.go
package db

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

// Connect opens a PostgreSQL connection.
func Connect(connStr string) (*sqlx.DB, error) {
    return sqlx.Connect("postgres", connStr)
}
