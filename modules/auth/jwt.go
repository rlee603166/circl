package auth

import (
    "os"
    "fmt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

type JWTService struct {
    secretKey []byte
}

func GetJWTService() *JWTService {
    secretKey := os.Getenv("JWT_SECRET")
    return &JWTService{secretKey: []byte(secretKey)}
}

func (s *JWTService) CreateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "username": username,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

    tokenString, err := token.SignedString(s.secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *JWTService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
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
