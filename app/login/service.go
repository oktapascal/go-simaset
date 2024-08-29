package login

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simaset/config"
	"github.com/oktapascal/go-simaset/exception"
	"github.com/oktapascal/go-simaset/helper"
	"github.com/oktapascal/go-simaset/model"
	"strings"
	"time"
)

type Service struct {
	rpo  model.LoginSessionRepository
	urpo model.UserRepository
	db   *sql.DB
}

func (svc *Service) Login(ctx context.Context, request *model.LoginRequest) model.LoginResponse {
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
		FlagCreate:  flagCreate,
		FlagRead:    flagRead,
		FlagUpdate:  flagUpdate,
		FlagDelete:  flagDelete,
		FlagApprove: flagApprove,
	}

	accessToken, errAccessToken := helper.GenerateAccessToken(&jwtParams)
	if errAccessToken != nil {
		panic(errAccessToken)
	}

	refreshToken, errRefreshToken := helper.GenerateRefreshToken(&jwtParams)
	if errRefreshToken != nil {
		panic(errRefreshToken)
	}

	session := new(model.LoginSession)
	session.UserId = user.Id
	session.RefreshToken = refreshToken

	expiresAt := helper.GetTime().Add(7 * (24 * time.Hour))
	session.ExpiresAt = expiresAt

	svc.rpo.CreateLoginSession(ctx, tx, session)

	return model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (svc *Service) Logout(ctx context.Context, claims jwt.MapClaims) {
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

	svc.rpo.RevokeLoginSession(ctx, tx, user.Id)
}

func (svc *Service) GenerateAccessToken(ctx context.Context, claims jwt.MapClaims) model.LoginResponse {
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

	session, errSession := svc.rpo.CheckRefreshToken(ctx, tx, user.Id)
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
		FlagCreate: flagCreate,
		FlagRead:   flagRead,
		FlagUpdate: flagUpdate,
		FlagDelete: flagDelete,
	}

	accessToken, errAccessToken := helper.GenerateAccessToken(&jwtParams)
	if errAccessToken != nil {
		panic(errAccessToken)
	}

	return model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
	}
}
