package login

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/config"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"strings"
	"time"
)

type Service struct {
	rpo  model.LoginSessionRepository
	urpo model.UserRepository
	db   *sql.DB
}

func (svc *Service) Login(ctx context.Context, request *model.LoginRequest, userAgent string) model.LoginResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	var identifier string
	if strings.Contains(request.Identifier, "@") {
		identifier = "email"
	} else {
		identifier = "username"
	}

	user := new(model.User)
	if identifier == "email" {
		user, err = svc.urpo.FindByEmail(ctx, tx, request.Identifier)
	} else {
		user, err = svc.urpo.FindByUsername(ctx, tx, request.Identifier)
	}

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	checkPassword := helper.CheckHash(request.Password, user.Password)
	if !checkPassword {
		panic(exception.NewNotMatchedError("password is not matched"))
	}

	listUserPermissions := svc.urpo.FindPermissionUser(ctx, tx, user.Id)
	var flagCreate, flagRead, flagUpdate, flagDelete, flagApprove bool
	for _, value := range *listUserPermissions {
		switch value.PermissionId {
		case "C":
			flagCreate = value.StatusPermit
		case "R":
			flagRead = value.StatusPermit
		case "U":
			flagUpdate = value.StatusPermit
		case "A":
			flagApprove = value.StatusPermit
		default:
			flagDelete = value.StatusPermit
		}
	}

	jwtParams := config.JwtParameters{
		Id:          user.Id,
		Email:       user.Email,
		Username:    user.Username,
		GroupMenu:   user.GroupMenu,
		FlagCreate:  flagCreate,
		FlagRead:    flagRead,
		FlagUpdate:  flagUpdate,
		FlagDelete:  flagDelete,
		FlagApprove: flagApprove,
	}

	accessToken, expiresAccessToken, errAccessToken := helper.GenerateAccessToken(&jwtParams)
	if errAccessToken != nil {
		panic(errAccessToken)
	}

	refreshToken, expiresRefreshToken, errRefreshToken := helper.GenerateRefreshToken(&jwtParams)
	if errRefreshToken != nil {
		panic(errRefreshToken)
	}

	session := new(model.LoginSession)
	session.UserId = user.Id
	session.RefreshToken = refreshToken
	session.UserAgent = userAgent

	unixFormat := expiresRefreshToken
	t := time.Unix(unixFormat, 0)
	formattedDateTime := t.Format("2006-01-02 15:04:05")
	expiresAt, _ := time.Parse("2006-01-02 15:04:05", formattedDateTime)
	session.ExpiresAt = expiresAt

	svc.rpo.CreateLoginSession(ctx, tx, session)

	userResponse := model.UserResponse{
		Username:    user.Username,
		Email:       user.Email,
		Name:        user.Name,
		Phone:       user.Phone,
		Photo:       user.Avatar,
		Permissions: *listUserPermissions,
	}

	accessTokenModel := model.AccessToken{
		Token:     accessToken,
		ExpiresAt: expiresAccessToken,
	}

	refreshTokenModel := model.RefreshToken{
		Token:     refreshToken,
		ExpiresAt: expiresRefreshToken,
	}

	return model.LoginResponse{
		AccessToken:  accessTokenModel,
		RefreshToken: refreshTokenModel,
		User:         userResponse,
	}
}

func (svc *Service) Logout(ctx context.Context, claims jwt.MapClaims, userAgent string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	username, ok := claims["sub"].(string)
	if !ok {
		panic("Something wrong when extracting username from jwt token")
	}

	user, errUser := svc.urpo.FindByUsername(ctx, tx, username)
	if errUser != nil {
		panic(exception.NewNotFoundError(errUser.Error()))
	}

	svc.rpo.RevokeLoginSession(ctx, tx, user.Id, userAgent)
}

func (svc *Service) GenerateAccessToken(ctx context.Context, claims jwt.MapClaims, userAgent string) model.TokenResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	username, ok := claims["sub"].(string)
	if !ok {
		panic("Something wrong when extracting username from jwt token")
	}

	user, errUser := svc.urpo.FindByUsername(ctx, tx, username)
	if errUser != nil {
		panic(exception.NewNotFoundError(errUser.Error()))
	}

	session, errSession := svc.rpo.CheckRefreshToken(ctx, tx, user.Id, userAgent)
	if errSession != nil {
		panic(exception.NewNotFoundError(errSession.Error()))
	}

	var flagCreate, flagRead, flagUpdate, flagDelete bool

	flagCreate, ok = claims["flag_create"].(bool)
	if !ok {
		panic("Something wrong when extracting flag create from jwt token")
	}

	flagRead, ok = claims["flag_read"].(bool)
	if !ok {
		panic("Something wrong when extracting flag read from jwt token")
	}

	flagUpdate, ok = claims["flag_update"].(bool)
	if !ok {
		panic("Something wrong when extracting flag update from jwt token")
	}

	flagDelete, ok = claims["flag_delete"].(bool)
	if !ok {
		panic("Something wrong when extracting flag delete from jwt token")
	}

	jwtParams := config.JwtParameters{
		Id:         user.Id,
		Email:      user.Email,
		Username:   user.Username,
		GroupMenu:  user.GroupMenu,
		FlagCreate: flagCreate,
		FlagRead:   flagRead,
		FlagUpdate: flagUpdate,
		FlagDelete: flagDelete,
	}

	accessToken, expiresAccessToken, errAccessToken := helper.GenerateAccessToken(&jwtParams)
	if errAccessToken != nil {
		panic(errAccessToken)
	}

	accessTokenModel := model.AccessToken{
		Token:     accessToken,
		ExpiresAt: expiresAccessToken,
	}

	refreshTokenModel := model.RefreshToken{
		Token:     session.RefreshToken,
		ExpiresAt: session.ExpiresAt.Unix(),
	}

	return model.TokenResponse{
		AccessToken:  accessTokenModel,
		RefreshToken: refreshTokenModel,
	}
}
