package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct {
}

func (rpo *Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error) {
	query := "select id, username, email, password, name, phone, avatar from users where email = ?"

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	user := new(model.User)
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.Avatar)

		if err != nil {
			panic(err)
		}

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (rpo *Repository) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*model.User, error) {
	query := "select id, username, email, password, name, phone, avatar from users where username = ?"

	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	user := new(model.User)
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.Avatar)

		if err != nil {
			panic(err)
		}

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (rpo *Repository) CreateUser(ctx context.Context, tx *sql.Tx, data *model.User) *model.User {
	query := `insert into users (id, username, email, password, name, phone, status_active, avatar) 
	values (UUID(),?,?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, data.Username, data.Email, data.Password, data.Name, data.Phone,
		data.StatusActive, data.Avatar)
	if err != nil {
		panic(err)
	}

	query = "select id from users order by created_at desc;"

	rows, errRows := tx.QueryContext(ctx, query)
	if errRows != nil {
		panic(errRows)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var id string
	if rows.Next() {
		err = rows.Scan(&id)
	}

	data.Id = id

	return data
}

func (rpo *Repository) CreateUserPermission(ctx context.Context, tx *sql.Tx, data *[]model.UserPermission) {
	placeholder := ""

	var args []any

	for i, row := range *data {
		placeholder += "(?, ?, ?)"

		if i < len(*data)-1 {
			placeholder += ","
		}

		args = append(args, row.UserId, row.PermissionId, row.StatusPermit)
	}

	query := fmt.Sprintf("insert into users_permissions (user_id, permission_id, status_permit) values %s", placeholder)

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) FindPermissionUser(ctx context.Context, tx *sql.Tx, userId string) *[]model.UserPermission {
	query := "select user_id, permission_id, status_permit from users_permissions where user_id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var userPermissions []model.UserPermission

	for rows.Next() {
		userPermission := model.UserPermission{}

		err = rows.Scan(&userPermission.UserId, &userPermission.PermissionId, &userPermission.StatusPermit)
		if err != nil {
			panic(err)
		}

		userPermissions = append(userPermissions, userPermission)
	}

	return &userPermissions
}
