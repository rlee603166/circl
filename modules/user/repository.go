// modules/user/repository.go
package user

import (
    "fmt"
    "github.com/jmoiron/sqlx"
)

// Repository encapsulates user CRUD operations.
type Repository struct {
    db *sqlx.DB
}

// NewRepository constructs a new user repository.
func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{db: db}
}

// Create inserts a new user record.
func (r *Repository) Create(u *User) (*User, error) {
    query := `INSERT INTO users (user_id, first_name, last_name, email, hashed_password, summary, pfp_url)
              VALUES (:user_id, :first_name, :last_name, :email, :hashed_password, :summary, :pfp_url)
              RETURNING user_id, first_name, last_name, email, hashed_password, summary, pfp_url`

    var created User
    stmt, err := r.db.PrepareNamed(query)
    if err != nil {
        return nil, err
    }

    err = stmt.Get(&created, u)
    if err != nil {
        return nil, err
    }

    return &created, nil
}

// GetByID fetches a user by their ID.
func (r *Repository) GetByID(id string) (*User, error) {
    var u User
    err := r.db.Get(&u, `SELECT * FROM users WHERE user_id=$1`, id)
    if err != nil {
        return nil, err
    }

    return &u, nil
}

// GetByEmail fetches a user by their Email.
func (r *Repository) GetByEmail(email string) (*User, error) {
    var u User
    err := r.db.Get(&u, `SELECT * FROM users WHERE email=$1`, email)
    if err != nil {
        fmt.Print(err)
        return nil, err
    }

    return &u, nil
}
