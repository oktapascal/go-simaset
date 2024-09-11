package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/config"
	"github.com/spf13/viper"
	"time"
)

func GenerateAccessToken(parameters *config.JwtParameters) (string, error) {
	token := config.GenerateToken(jwt.MapClaims{
		"iss":          viper.GetString("APP_NAME"),
		"sub":          parameters.Username,
		"aud":          parameters.Email,
		"exp":          GetTime().Add(15 * time.Minute).Unix(),
		"iat":          GetTime().Unix(),
		"id":           parameters.Id,
		"flag_create":  parameters.FlagCreate,
		"flag_read":    parameters.FlagRead,
		"flag_update":  parameters.FlagUpdate,
		"flag_delete":  parameters.FlagDelete,
		"flag_approve": parameters.FlagApprove,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SIGNATURE_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(parameters *config.JwtParameters) (string, error) {
	token := config.GenerateToken(jwt.MapClaims{
		"iss":         viper.GetString("APP_NAME"),
		"sub":         parameters.Username,
		"aud":         parameters.Email,
		"exp":         GetTime().Add(7 * (24 * time.Hour)).Unix(),
		"iat":         GetTime().Unix(),
		"flag_create": parameters.FlagCreate,
		"flag_read":   parameters.FlagRead,
		"flag_update": parameters.FlagUpdate,
		"flag_delete": parameters.FlagDelete,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_REFRESH_SIGNATURE_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := config.VerifyToken(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(viper.GetString("JWT_SIGNATURE_KEY")), nil
	})

	return token, err
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := config.VerifyToken(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(viper.GetString("JWT_REFRESH_SIGNATURE_KEY")), nil
	})

	return token, err
}
