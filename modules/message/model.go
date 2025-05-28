package message

type Message struct {
    MessageID   int `json:"message_id" db:"message_id"`
    SessionID   string `json:"session_id" db:"session_id"`
    Role        string `json:"role" db:"role"`
    Content     string `json:"content" db:"content"`
    CreatedAt   string `json:"created_at" db:"created_at"`
}

type CreateMessage struct {
    SessionID   string `json:"session_id" db:"session_id"`
    Role        string `json:"role" db:"role"`
    Content     string `json:"content" db:"content"`
}
