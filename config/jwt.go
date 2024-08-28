package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JwtParameters struct {
	Id         string
	Email      string
	Username   string
	FlagCreate bool
	FlagRead   bool
	FlagUpdate bool
	FlagDelete bool
}

func GenerateToken(claims jwt.MapClaims) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}

func VerifyToken(tokenString string, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
