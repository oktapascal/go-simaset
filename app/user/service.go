package user

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simaset/exception"
	"github.com/oktapascal/go-simaset/helper"
	"github.com/oktapascal/go-simaset/model"
)

type Service struct {
	rpo model.UserRepository
	db  *sql.DB
}

func (svc *Service) SaveUser(ctx context.Context, request *model.SaveUserRequest) model.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	_, err = svc.rpo.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		panic(exception.NewDuplicateError("email already exists"))
	}

	_, err = svc.rpo.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(exception.NewDuplicateError("username already exists"))
	}

	user := new(model.User)
	user.Email = request.Email
	user.Username = request.Username
	user.Phone = request.Phone
	user.Name = request.Name
	user.StatusActive = true

	hash, errHash := helper.Hash(request.Password)
	if errHash == nil {
		user.Password = hash
	}

	user = svc.rpo.CreateUser(ctx, tx, user)

	var permissions []model.UserPermission

	for _, value := range request.Permissions {
		permission := model.UserPermission{
			UserId:       user.Id,
			PermissionId: value.PermissionId,
			StatusPermit: *value.StatusPermit,
		}

		permissions = append(permissions, permission)
	}

	svc.rpo.CreateUserPermission(ctx, tx, &permissions)

	return model.UserResponse{
		Username:    user.Username,
		Email:       user.Email,
		Name:        user.Name,
		Phone:       user.Phone,
		Permissions: permissions,
	}
}

func (svc *Service) GetUserByToken(ctx context.Context, claims jwt.MapClaims) model.UserProfileResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	username, ok := claims["sub"].(string)
	if !ok {
		panic("Something wrong when extracting username from jwt token")
	}

	user, errUser := svc.rpo.FindByUsername(ctx, tx, username)
	if errUser != nil {
		panic(exception.NewNotFoundError(errUser.Error()))
	}

	return model.UserProfileResponse{
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Phone:    user.Phone,
	}
}
