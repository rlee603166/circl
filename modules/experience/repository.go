// modules/experience/repository.go
package experience

import "github.com/jmoiron/sqlx"

type Repository struct { db *sqlx.DB }

func NewRepository(db *sqlx.DB) *Repository { return &Repository{db} }

func (r *Repository) Create(e *Experience) error {
    q := `INSERT INTO experiences (user_id, company_name, job_title, location, start_date, end_date, experience_description)
          VALUES (:user_id, :company_name, :job_title, :location, :start_date, :end_date, :experience_description)`
    _, err := r.db.NamedExec(q, e)
    return err
}

func (r *Repository) GetByUserID(userID string) ([]Experience, error) {
    var list []Experience
    err := r.db.Select(&list, `SELECT * FROM experiences WHERE user_id=$1`, userID)
    return list, err
}

func (r *Repository) Update(e *Experience) error {
    q := `UPDATE experiences SET company_name=:company_name, job_title=:job_title, location=:location,
          start_date=:start_date, end_date=:end_date, experience_description=:experience_description
          WHERE experience_id=:experience_id`            
    _, err := r.db.NamedExec(q, e)
    return err
}

func (r *Repository) Delete(id int) error {
    _, err := r.db.Exec(`DELETE FROM experiences WHERE experience_id=$1`, id)
    return err
}
