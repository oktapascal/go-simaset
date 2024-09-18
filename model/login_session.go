package model

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type (
	LoginSession struct {
		Id           string
		UserId       string
		RefreshToken string
		Revoked      bool
		UserAgent    string
		ExpiresAt    time.Time
	}

	LoginRequest struct {
		Identifier string `validate:"required" json:"identifier"`
		Password   string `validate:"required" json:"password"`
	}

	AccessToken struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}

	RefreshToken struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}

	LoginResponse struct {
		AccessToken  AccessToken  `json:"access_token"`
		RefreshToken RefreshToken `json:"refresh_token"`
		User         UserResponse `json:"user"`
	}

	TokenResponse struct {
		AccessToken  AccessToken  `json:"access_token"`
		RefreshToken RefreshToken `json:"refresh_token"`
	}

	LoginSessionRepository interface {
		CreateLoginSession(ctx context.Context, tx *sql.Tx, data *LoginSession)
		RevokeLoginSession(ctx context.Context, tx *sql.Tx, userId string, userAgent string)
		CheckRefreshToken(ctx context.Context, tx *sql.Tx, userId string, userAgent string) (*LoginSession, error)
	}

	LoginSessionService interface {
		Login(ctx context.Context, request *LoginRequest, userAgent string) LoginResponse
		Logout(ctx context.Context, claims jwt.MapClaims, userAgen string)
		GenerateAccessToken(ctx context.Context, claims jwt.MapClaims, userAgent string) TokenResponse
	}

	LoginSessionHandler interface {
		Login() http.HandlerFunc
		Logout() http.HandlerFunc
		GetAccessToken() http.HandlerFunc
	}
)
