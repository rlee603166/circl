// modules/user/service.go
package user

import (
    "errors"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

// Service holds business logic for users.
type Service struct {
    repo *Repository
}

// NewService creates a new User service.
func NewService(r *Repository) *Service {
    return &Service{repo: r}
}

// CreateUser validates input, hashes the password, and stores the user.
func (s *Service) CreateUser(u *User) (*User, error) {
    if u.Email == nil || u.HashedPassword == nil || *u.Email == "" || *u.HashedPassword == "" {
        return nil, errors.New("email and password are required")
    }

    u.UserID = uuid.New().String()

    hash, err := bcrypt.GenerateFromPassword([]byte(*u.HashedPassword), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    hashed := string(hash)
    u.HashedPassword = &hashed

    if err := s.repo.Create(u); err != nil {
        return nil, err
    }

    return u, nil
}

// CreateUser validates input, hashes the password, and stores the user.
func (s *Service) CreateGoogleUser(u *User) (*User, error) {
    if u.Email == nil ||*u.Email == "" {
        return nil, errors.New("email and password are required")
    }

    u.UserID = uuid.New().String()

    if err := s.repo.Create(u); err != nil {
        return nil, err
    }

    return u, nil
}

// GetUserByID retrieves a user by ID.
func (s *Service) GetUserByID(id string) (*User, error) {
    return s.repo.GetByID(id)
}

// GetUserByEmail retrieves a user by Email.
func (s *Service) GetUserByEmail(email string) (*User, error) {
    return s.repo.GetByEmail(email)
}
