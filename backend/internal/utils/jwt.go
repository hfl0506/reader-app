package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtPayload struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type JwtCustomClaims struct {
	*JwtPayload
	jwt.RegisteredClaims
}

func GenerateToken(claims *JwtPayload, expire time.Duration) (string, error) {
	tokenClaims := &JwtCustomClaims{
		claims,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	signingKey := []byte("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	t, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, errors.New("Parse token string failed")
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("User invalid")
	}
}
