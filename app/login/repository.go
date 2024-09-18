package login

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct {
}

func (rpo *Repository) CreateLoginSession(ctx context.Context, tx *sql.Tx, data *model.LoginSession) {
	query := `insert into login_sessions (id, user_id, refresh_token, user_agent, revoked, expired_at)
	values (UUID(), ?, ?, ?, false, ?)`

	_, err := tx.ExecContext(ctx, query, data.UserId, data.RefreshToken, data.UserAgent, data.ExpiresAt)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) RevokeLoginSession(ctx context.Context, tx *sql.Tx, userId string, userAgent string) {
	query := "update login_sessions set revoked = true where user_id = ? and user_agent = ?"

	_, err := tx.ExecContext(ctx, query, userId, userAgent)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) CheckRefreshToken(ctx context.Context, tx *sql.Tx, userId string, userAgent string) (*model.LoginSession, error) {
	query := "select id, user_id, refresh_token, expired_at from login_sessions where user_id = ? and user_agent = ? and revoked = false"

	rows, err := tx.QueryContext(ctx, query, userId, userAgent)
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
