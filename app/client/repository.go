package client

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) CreateClient(ctx context.Context, tx *sql.Tx, data *model.Client) *model.Client {
	query := "insert into clients (id, name, address, phone) values (UUID(), ?, ?, ?)"

	_, err := tx.Exec(query, data.Id, data.Name, data.Address, data.Phone)
	if err != nil {
		panic(err)
	}

	query = "select id from clients order by created_at desc"

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

func (rpo *Repository) CreateClientPic(ctx context.Context, tx *sql.Tx, data *[]model.ClientPic) *[]model.ClientPic {
	placeholder := ""

	var args []any

	for i, row := range *data {
		placeholder += "(UUID(), ?, ?, ?, ?, ?)"

		if i < len(*data)-1 {
			placeholder += ","
		}

		args = append(args, row.ClientId, row.Name, row.Phone, row.Email, row.Address)
	}

	query := fmt.Sprintf("insert into clients_pic (id, client_id, name, phone, email, address) values %s", placeholder)

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}

	return data
}
