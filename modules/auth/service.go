package auth

import (
    "os"
    "fmt"
    "time"
    "context"
 
    "github.com/golang-jwt/jwt/v5"
    "google.golang.org/api/idtoken"
)

type Service struct {
    secretKey []byte
    googleClientID string
}

func GetService() *Service {
    secretKey := os.Getenv("JWT_SECRET")
    googleClientID := os.Getenv("GOOGLE_ClIENT_ID")
    return &Service{
        secretKey:      []byte(secretKey),
        googleClientID: googleClientID,
    }
}

func (s *Service) CreateToken(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "email": email,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

    tokenString, err := token.SignedString(s.secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *Service) VerifyToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return s.secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}

func (s *Service) VerifyGoogleToken(tokenString string) (*AuthPayload, error) {
    payload, err := idtoken.Validate(context.Background(), tokenString, s.googleClientID)
    if err != nil {
        return nil, fmt.Errorf("token validation failed: %w", err)
    }

    claims := payload.Claims

    return &AuthPayload{
        FirstName:  getStringClaim(claims, "given_name"),
        LastName:   getStringClaim(claims, "family_name"),
        Email:      getStringClaim(claims, "email"),
    }, nil
}

func getStringClaim(claims map[string]interface{}, key string) string {
    if val, ok := claims[key]; ok {
        if str, ok := val.(string); ok {
            return str
        }
    }

    return ""
}
