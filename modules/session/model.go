package session

type Session struct {
    SessionID   string `json:"session_id" db:"session_id"`
    UserID      string `json:"user_id" db:"user_id"`
    CreatedAt   string `json:"created_at" db:"created_at"`
}

type CreateSession struct {
    SessionID   string `json:"session_id" db:"session_id"`
    UserID      string `json:"user_id" db:"user_id"`
}
