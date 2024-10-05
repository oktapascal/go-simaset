package client

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/model"
	"strconv"
	"strings"
)

type Repository struct{}

func (rpo *Repository) GenerateClientKode(ctx context.Context, tx *sql.Tx) *string {
	query := "select id from clients order by created_at desc limit 1"

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

	var id string
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}

		strNumber := id[4:]
		number, errConvert := strconv.Atoi(strNumber)
		if errConvert != nil {
			panic(errConvert)
		}

		number++
		strNumber = strconv.Itoa(number)

		if len(strNumber) == 3 {
			id = fmt.Sprintf("KTG-%s", strNumber)
		} else if len(strNumber) == 2 {
			id = fmt.Sprintf("KTG-0%s", strNumber)
		} else {
			id = fmt.Sprintf("KTG-00%s", strNumber)
		}
	} else {
		id = "KTG-001"
	}

	return &id
}

func (rpo *Repository) CreateClient(ctx context.Context, tx *sql.Tx, data *model.Client) *model.Client {
	query := "insert into clients (id, name, address, phone) values (?, ?, ?, ?)"

	_, err := tx.ExecContext(ctx, query, data.Id, data.Name, data.Address, data.Phone)
	if err != nil {
		panic(err)
	}

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
	query := `select t1.id, t1.name, t1.address, t1.phone, t2.jumlah_pic
	from clients t1
	inner join (
		select client_id, count(id) jumlah_pic
		from clients_pic
		where deleted_at is null
		group by client_id
	) t2 on t1.id=t2.client_id
	where t1.deleted_at is null
	order by t1.created_at desc, t1.updated_at desc`

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
		err = rows.Scan(&client.Id, &client.Name, &client.Address, &client.Phone, &client.NumberOfPics)
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
	query := "select id, name, phone, email, address from clients_pic where client_id = ? and deleted_at is null"

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

func (rpo *Repository) UpdateClient(ctx context.Context, tx *sql.Tx, data *model.Client) *model.Client {
	query := "update clients set name = ?, address = ?, phone = ? where id = ?"

	_, err := tx.ExecContext(ctx, query, data.Name, data.Address, data.Phone, data.Id)
	if err != nil {
		panic(err)
	}

	return data
}

func (rpo *Repository) UpdateClientPic(ctx context.Context, tx *sql.Tx, data *[]model.ClientPic) *[]model.ClientPic {
	query := "update clients_pic set name = ?, phone = ?, email = ?, address = ? where id = ? and client_id = ?"

	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	var updates []struct {
		Name     string
		Phone    string
		Email    string
		Address  string
		Id       string
		ClientId string
	}

	for _, value := range *data {
		updates = append(updates, struct {
			Name     string
			Phone    string
			Email    string
			Address  string
			Id       string
			ClientId string
		}{Name: value.Name, Phone: value.Phone, Email: value.Email, Address: value.Address, Id: value.Id, ClientId: value.ClientId})
	}

	for _, update := range updates {
		_, err := stmt.ExecContext(ctx, update.Name, update.Phone, update.Email, update.Address, update.Id, update.ClientId)
		if err != nil {
			panic(err)
		}
	}

	return data
}

func (rpo *Repository) DeleteClientPic(ctx context.Context, tx *sql.Tx, id string, clientId []string) {
	placeholders := make([]string, len(clientId))
	for i := range clientId {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf("update clients_pic set deleted_at = current_timestamp where client_id = ? and id not in (%s)", strings.Join(placeholders, ","))

	args := make([]any, len(clientId)+1)
	args[0] = id
	for i, value := range clientId {
		args[i+1] = value
	}

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) DeleteClient(ctx context.Context, tx *sql.Tx, id string) {
	query := "update clients set deleted_at = current_timestamp where id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	query = "update clients_pic set deleted_at = current_timestamp where client_id = ? and deleted_at is null"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}
