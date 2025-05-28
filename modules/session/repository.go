package session

import (
    "github.com/jmoiron/sqlx"
)

type Repository struct {
    db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(s *CreateSession) (*Session, error) {
    query := `INSERT INTO chat_sessions (session_id, user_id)
              VALUES (:session_id, :user_id)
              RETURNING session_id, user_id, created_at`

    var created Session
    stmt, err := r.db.PrepareNamed(query)
    if err != nil {
        return nil, err
    }

    err = stmt.Get(&created, s)
    if err != nil {
        return nil, err
    }

    return &created, nil
}

func (r *Repository) GetByID(id string) (*Session, error) {
    var s Session
    err := r.db.Get(&s, `SELECT * FROM chat_sessions WHERE session_id=$1`, id)
    if err != nil {
        return nil, err
    }
     
    return &s, nil
}

func (r *Repository) GetByUserID(userID string) ([]Session, error) {
    var list []Session
    err := r.db.Select(&list, `SELECT * FROM chat_sessions WHERE user_id=$1`, userID)
    return list, err
}
