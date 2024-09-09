package model

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type (
	User struct {
		Id           string
		Username     string
		Email        string
		Password     string
		Name         string
		Phone        string
		Avatar       string
		StatusActive bool
		DeletedAt    time.Time
	}

	UserPermission struct {
		UserId       string `json:"user_id"`
		PermissionId string `json:"permission_id"`
		StatusPermit bool   `json:"status_permit"`
	}

	userPermissionRequest struct {
		PermissionId string `validate:"required,oneof=C R U D A" json:"permission_id"`
		StatusPermit *bool  `validate:"required" json:"status_permit"`
	}

	SaveUserRequest struct {
		Username             string                  `json:"username" validate:"required,min=1,max=50"`
		Email                string                  `json:"email" validate:"required,email,min=1,max=50"`
		Password             string                  `json:"password" validate:"required,min=1,max=50"`
		PasswordConfirmation string                  `json:"password_confirmation" validate:"required,eqfield=Password"`
		Name                 string                  `json:"name" validate:"required,min=1,max=50"`
		Phone                string                  `json:"phone" validate:"required,min=11,max=13"`
		Permissions          []userPermissionRequest `json:"permissions" validate:"dive"`
	}

	UserResponse struct {
		Username    string           `json:"username"`
		Email       string           `json:"email"`
		Name        string           `json:"name"`
		Phone       string           `json:"phone"`
		Permissions []UserPermission `json:"permissions"`
	}

	UserProfileResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
	}

	UserRepository interface {
		FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error)
		FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*User, error)
		CreateUser(ctx context.Context, tx *sql.Tx, data *User) *User
		CreateUserPermission(ctx context.Context, tx *sql.Tx, data *[]UserPermission)
		FindPermissionUser(ctx context.Context, tx *sql.Tx, userId string) *[]UserPermission
	}

	UserService interface {
		SaveUser(ctx context.Context, request *SaveUserRequest) UserResponse
		GetUserByToken(ctx context.Context, claims jwt.MapClaims) UserProfileResponse
	}

	UserHandler interface {
		SaveUser() http.HandlerFunc
		GetUserByToken() http.HandlerFunc
	}
)
