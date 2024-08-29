package login

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simaset/model"
)

type Repository struct {
}

func (rpo *Repository) CreateLoginSession(ctx context.Context, tx *sql.Tx, data *model.LoginSession) {
	query := "insert into login_sessions (id, user_id, refresh_token, revoked, expired_at) values (UUID(),?,?,false,?)"

	_, err := tx.ExecContext(ctx, query, data.UserId, data.RefreshToken, data.ExpiresAt)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) RevokeLoginSession(ctx context.Context, tx *sql.Tx, userId string) {
	query := "update login_sessions set revoked = true where user_id = ?"

	_, err := tx.ExecContext(ctx, query, userId)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) CheckRefreshToken(ctx context.Context, tx *sql.Tx, userId string) (*model.LoginSession, error) {
	query := "select id, user_id, refresh_token, expired_at from login_sessions where user_id = ? and revoked = false"

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

	session := new(model.LoginSession)
	if rows.Next() {
		err = rows.Scan(&session.Id, &session.UserId, &session.RefreshToken, &session.ExpiresAt)
		if err != nil {
			panic(err)
		}

		return session, nil
	} else {
		return session, errors.New("session not found")
	}
}
