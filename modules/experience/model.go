// modules/experience/model.go
package experience

type Experience struct {
    ExperienceID          int     `json:"experience_id" db:"experience_id"`
    UserID                string  `json:"user_id" db:"user_id"`
    CompanyName           string  `json:"company_name" db:"company_name"`
    JobTitle              string  `json:"job_title" db:"job_title"`
    Location              *string `json:"location,omitempty" db:"location"`
    StartDate             *string  `json:"start_date" db:"start_date"`
    EndDate               *string `json:"end_date,omitempty" db:"end_date"`
    ExperienceDescription *string `json:"experience_description,omitempty" db:"experience_description"`
}

type CreateExperience struct {
    UserID                string  `json:"user_id" db:"user_id"`
    CompanyName           string  `json:"company_name" db:"company_name"`
    JobTitle              string  `json:"job_title" db:"job_title"`
    Location              *string `json:"location,omitempty" db:"location"`
    StartDate             *string  `json:"start_date" db:"start_date"`
    EndDate               *string `json:"end_date,omitempty" db:"end_date"`
    ExperienceDescription *string `json:"experience_description,omitempty" db:"experience_description"`
}
