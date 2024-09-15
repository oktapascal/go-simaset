package client

import (
	"database/sql"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.ClientRepository
	db  *sql.DB
}
