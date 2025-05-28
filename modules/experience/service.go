// modules/experience/service.go
package experience

import "errors"

type Service struct { repo *Repository }

func NewService(r *Repository) *Service { return &Service{r} }

func (s *Service) CreateExperience(e *CreateExperience) (*Experience, error) {
    if e.UserID == "" { return nil, errors.New("user_id required") }
    created, err := s.repo.Create(e)
    if err != nil { return nil, err }
    return created, nil
}

func (s *Service) GetExperiencesByUserID(userID string) ([]Experience, error) {
    return s.repo.GetByUserID(userID)
}

func (s *Service) UpdateExperience(e *Experience) (*Experience, error) {
    if err := s.repo.Update(e); err != nil { return nil, err }
    return e, nil
}

func (s *Service) DeleteExperience(id int) error {
    return s.repo.Delete(id)
}
