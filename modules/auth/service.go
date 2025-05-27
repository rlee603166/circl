package auth

import (
    "os"
    "fmt"
    "time"
    "context"
 
    "github.com/golang-jwt/jwt/v5"
    "google.golang.org/api/idtoken"
    "github.com/rlee603166/circl/modules/user"
)

type Service struct {
    secretAccessKey     []byte
    secretRefreshKey    []byte
    googleClientID      string
}

func GetService() *Service {
    secretAccessKey     := os.Getenv("JWT_ACCESS_SECRET")
    secretRefreshKey    := os.Getenv("JWT_REFRESH_SECRET")
    googleClientID      := os.Getenv("GOOGLE_ClIENT_ID")

    return &Service{
        secretAccessKey:    []byte(secretAccessKey),
        secretRefreshKey:   []byte(secretRefreshKey),
        googleClientID:     googleClientID,
    }
}

func (s *Service) CreateAccessToken(u *user.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "userID": u.UserID,
            "email": u.Email,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

    tokenString, err := token.SignedString(s.secretAccessKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *Service) CreateRefreshToken(u *user.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "userID": u.UserID,
            "email": u.Email,
            "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
        })

    tokenString, err := token.SignedString(s.secretRefreshKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *Service) VerifyAccessToken(tokenString string) (*TokenPayload, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return s.secretAccessKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    expRaw, ok := claims["exp"]
    if !ok {
        return nil, fmt.Errorf("missing exp claim")
    }

    expFloat, ok := expRaw.(float64)
    if !ok {
        return nil, fmt.Errorf("invalid exp claim")
    }

    if int64(expFloat) < time.Now().Unix() {
        return nil, fmt.Errorf("token expired")

    }

    return &TokenPayload{
        UserID: getStringClaim(claims, "userID"),
        Email:  getStringClaim(claims, "email"),
    }, nil
}

func (s *Service) VerifyRefreshToken(tokenString string) (*TokenPayload, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return s.secretRefreshKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    expRaw, ok := claims["exp"]
    if !ok {
        return nil, fmt.Errorf("missing exp claim")
    }

    expFloat, ok := expRaw.(float64)
    if !ok {
        return nil, fmt.Errorf("invalid exp claim")
    }

    if int64(expFloat) < time.Now().Unix() {
        return nil, fmt.Errorf("token expired")

    }

    return &TokenPayload{
        UserID: getStringClaim(claims, "userID"),
        Email:  getStringClaim(claims, "email"),
    }, nil
}

func (s *Service) VerifyGoogleToken(tokenString string) (*GooglePayload, error) {
    payload, err := idtoken.Validate(context.Background(), tokenString, s.googleClientID)
    if err != nil {
        return nil, fmt.Errorf("token validation failed: %w", err)
    }

    claims := payload.Claims

    return &GooglePayload{
        FirstName:  getStringClaim(claims, "given_name"),
        LastName:   getStringClaim(claims, "family_name"),
        Email:      getStringClaim(claims, "email"),
        PfpURL:     getStringClaim(claims, "picture"),
    }, nil
}

func getStringClaim(claims jwt.MapClaims, key string) *string {
    if val, ok := claims[key]; ok {
        if str, ok := val.(string); ok {
            return &str
        }
    }
    return nil
}
