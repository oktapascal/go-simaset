package client

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) CreateClient(ctx context.Context, tx *sql.Tx, data *model.Client) *model.Client {
	query := "insert into clients (id, name, address, phone) values (UUID(), ?, ?, ?)"

	_, err := tx.Exec(query, data.Name, data.Address, data.Phone)
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

func (rpo *Repository) GetAllClients(ctx context.Context, tx *sql.Tx) *[]model.Client {
	query := "select id, name, address, phone from clients where deleted_at is null"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var clients []model.Client
	for rows.Next() {
		var client model.Client
		err = rows.Scan(&client.Id, &client.Name, &client.Address, &client.Phone)
		if err != nil {
			panic(err)
		}

		clients = append(clients, client)
	}

	return &clients
}

func (rpo *Repository) GetClient(ctx context.Context, tx *sql.Tx, id string) (*model.Client, error) {
	query := "select id, name, phone, address from clients where id = ?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	client := new(model.Client)
	if rows.Next() {
		err = rows.Scan(&client.Id, &client.Name, &client.Phone, &client.Address)
		if err != nil {
			panic(err)
		}

		return client, nil
	} else {
		return nil, errors.New("client not found")
	}
}

func (rpo *Repository) GetClientPic(ctx context.Context, tx *sql.Tx, id string) *[]model.ClientPic {
	query := "select id, name, phone, email, address from clients_pic where client_id = ?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var clientsPic []model.ClientPic
	for rows.Next() {
		var clientPic model.ClientPic
		err = rows.Scan(&clientPic.Id, &clientPic.Name, &clientPic.Phone, &clientPic.Email, &clientPic.Address)
		if err != nil {
			panic(err)
		}

		clientsPic = append(clientsPic, clientPic)
	}

	return &clientsPic
}
