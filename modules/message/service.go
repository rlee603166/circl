package message

import (
    "errors"
)

type Service struct {
    repo *Repository
}

func NewService(r *Repository) *Service {
    return &Service{repo: r}
}

func (s *Service) CreateMessage(m *CreateMessage) (*Message, error) {
    if m.SessionID == "" || m.Role == "" || m.Content == "" {
        return nil, errors.New("sessionid, role, and content are required")
    }

    created, err := s.repo.Create(m)
    if err != nil {
        return nil, err 
    }

    return created, nil 
}

func (s *Service) GetMessageByID(id string) (*Message, error) {
    return s.repo.GetByID(id)
}

func (s *Service) GetMessagesByUserID(userID string) ([]Message, error) {
    return s.repo.GetByUserID(userID)
}

func (s *Service) GetMessagesBySessionID(sessionID string) ([]Message, error) {
    return s.repo.GetBySessionID(sessionID)
}
