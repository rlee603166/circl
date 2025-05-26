package auth

type GooglePayload struct {
    FirstName   string
    LastName    string
    Username    string
    Email       string
}

type TokenPayload struct {
    UserID  string
    Email   string
}
