package session

import (
    "errors"
)

type Service struct {
    repo *Repository
}

func NewService(r *Repository) *Service {
    return &Service{repo: r}
}

func (svc *Service) CreateSession(s *CreateSession) (*Session, error) {
    // if s.UserID == nil || *s.UserID == nil {
    if s.UserID == "" || s.SessionID == "" {
        return nil, errors.New("userID and sessionID are required")
    }

    created, err := svc.repo.Create(s)
    if err != nil {
        return nil, err
    }

    return created, nil
}

func (svc *Service) GetSessionByID(id string) (*Session, error) {
    return svc.repo.GetByID(id)
}

func (svc *Service) GetSessionsByUserID(userID string) ([]Session, error) {
    return svc.repo.GetByUserID(userID)
}
