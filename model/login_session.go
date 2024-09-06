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
		ExpiresAt    time.Time
	}

	LoginRequest struct {
		Identifier string `validate:"required" json:"identifier"`
		Password   string `validate:"required" json:"password"`
	}

	LoginResponse struct {
		AccessToken  string       `json:"access_token"`
		RefreshToken string       `json:"refresh_token"`
		User         UserResponse `json:"user"`
	}

	LoginSessionRepository interface {
		CreateLoginSession(ctx context.Context, tx *sql.Tx, data *LoginSession)
		RevokeLoginSession(ctx context.Context, tx *sql.Tx, userId string)
		CheckRefreshToken(ctx context.Context, tx *sql.Tx, userId string) (*LoginSession, error)
	}

	LoginSessionService interface {
		Login(ctx context.Context, request *LoginRequest) LoginResponse
		Logout(ctx context.Context, claims jwt.MapClaims)
		GenerateAccessToken(ctx context.Context, claims jwt.MapClaims) LoginResponse
	}

	LoginSessionHandler interface {
		Login() http.HandlerFunc
		Logout() http.HandlerFunc
		GetAccessToken() http.HandlerFunc
	}
)
