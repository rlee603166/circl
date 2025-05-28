// modules/education/repository.go
package education

import "github.com/jmoiron/sqlx"

type Repository struct { db *sqlx.DB }

func NewRepository(db *sqlx.DB) *Repository { return &Repository{db} }

func (r *Repository) Create(e *CreateEducation) (*Education, error) {
    query := `INSERT INTO educations (user_id, institution_name, degree_type, degree_name, enrollment_date, graduation_date)
          VALUES (:user_id, :institution_name, :degree_type, :degree_name, :enrollment_date, :graduation_date)
          RETURNING education_id, user_id, institution_name, degree_type, degree_name, enrollment_date, graduation_date`

    var created Education
    stmt, err := r.db.PrepareNamed(query)
    if err != nil {
        return nil, err
    }

    err = stmt.Get(&created, e)
    if err != nil {
        return nil, err
    }

    return &created, nil
}

func (r *Repository) GetByUserID(userID string) ([]Education, error) {
    var list []Education
    err := r.db.Select(&list, `SELECT * FROM educations WHERE user_id=$1`, userID)
    return list, err
}

func (r *Repository) Update(e *Education) error {
    q := `UPDATE educations SET institution_name=:institution_name, degree_type=:degree_type,
          degree_name=:degree_name, enrollment_date=:enrollment_date, graduation_date=:graduation_date
          WHERE education_id=:education_id`
    _, err := r.db.NamedExec(q, e)
    return err
}

func (r *Repository) Delete(id int) error {
    _, err := r.db.Exec(`DELETE FROM educations WHERE education_id=$1`, id)
    return err
}
