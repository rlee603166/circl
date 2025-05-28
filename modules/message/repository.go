package message

import (
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

func (r *Repository) Create(m *CreateMessage) (*Message, error) {
    query := `INSERT INTO chat_messages (session_id, role, content)
              VALUES (:message_id, :session_id, :role, :content)
              RETURNING message_id, session_id, session_id, role, content, created_at`

    var created Message
    stmt, err := r.db.PrepareNamed(query)
    if err != nil {
        return nil, err
    }

    err = stmt.Get(&created, m)
    if err != nil {
        return nil, err
    }

    return &created, nil
}

func (r *Repository) GetByID(id string) (*Message, error) {
    var m Message 
    err := r.db.Get(&m, `SELECT * FROM chat_messages WHERE message_id=$1`, id)
    if err != nil {
        return nil, err
    }
     
    return &m, nil
}

func (r *Repository) GetByUserID(userID string) ([]Message, error) {
    var list []Message
    err := r.db.Select(&list, `SELECT * FROM chat_messages WHERE user_id=$1`, userID)
    return list, err
}

func (r *Repository) GetBySessionID(sessionID string) ([]Message, error) {
    var list []Message
    err := r.db.Select(&list, `SELECT * FROM chat_messages WHERE session_id=$1`, sessionID)
    return list, err
}
