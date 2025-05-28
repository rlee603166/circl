// modules/education/model.go
package education

type Education struct {
    EducationID     int     `json:"education_id" db:"education_id"`
    UserID          string  `json:"user_id" db:"user_id"`
    InstitutionName string  `json:"institution_name" db:"institution_name"`
    DegreeType      *string `json:"degree_type,omitempty" db:"degree_type"`
    DegreeName      *string `json:"degree_name,omitempty" db:"degree_name"`
    EnrollmentDate  *string  `json:"enrollment_date" db:"enrollment_date"`
    GraduationDate  *string `json:"graduation_date,omitempty" db:"graduation_date"`
}

type CreateEducation struct {
    UserID          string  `json:"user_id" db:"user_id"`
    InstitutionName string  `json:"institution_name" db:"institution_name"`
    DegreeType      *string `json:"degree_type,omitempty" db:"degree_type"`
    DegreeName      *string `json:"degree_name,omitempty" db:"degree_name"`
    EnrollmentDate  *string  `json:"enrollment_date" db:"enrollment_date"`
    GraduationDate  *string `json:"graduation_date,omitempty" db:"graduation_date"`
}

