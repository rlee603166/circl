// modules/education/service.go
package education

import "errors"

type Service struct { repo *Repository }

func NewService(r *Repository) *Service { return &Service{r} }

func (s *Service) CreateEducation(e *CreateEducation) (*Education, error) {
    if e.UserID == "" { return nil, errors.New("user_id required") }
    created, err := s.repo.Create(e); 
    if err != nil { 
        return nil, err 
    }
    return created, nil
}

func (s *Service) GetEducationsByUserID(userID string) ([]Education, error) {
    return s.repo.GetByUserID(userID)
}

func (s *Service) UpdateEducation(e *Education) (*Education, error) {
    if err := s.repo.Update(e); err != nil { return nil, err }
    return e, nil
}

func (s *Service) DeleteEducation(id int) error {
    return s.repo.Delete(id)
}
