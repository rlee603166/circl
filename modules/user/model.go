package user

type User struct {
    UserID         string `json:"user_id" db:"user_id"`
    FirstName      string `json:"first_name" db:"first_name"`
    LastName       string `json:"last_name" db:"last_name"`
    Email          string `json:"email" db:"email"`
    HashedPassword string `json:"hashed_password" db:"hashed_password"`
    Summary        string `json:"summary" db:"summary"`
    PfpURL         string `json:"pfp_url" db:"pfp_url"`
}
