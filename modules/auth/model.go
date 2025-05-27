package auth

type GooglePayload struct {
    FirstName   *string
    LastName    *string
    Email       *string
    PfpURL      *string
}

type TokenPayload struct {
    UserID  *string
    Email   *string
}
