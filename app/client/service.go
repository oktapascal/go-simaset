package client

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.ClientRepository
	db  *sql.DB
}

func (svc *Service) StoreClient(ctx context.Context, request *model.SaveClientRequest) model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	client := new(model.Client)
	client.Name = request.Name
	client.Phone = request.Phone
	client.Address = request.Address

	client = svc.rpo.CreateClient(ctx, tx, client)

	var clientsPic []model.ClientPic

	for _, value := range request.ClientPic {
		clientPic := model.ClientPic{
			ClientId: client.Id,
			Name:     value.Name,
			Phone:    value.Phone,
			Email:    value.Email,
			Address:  value.Address,
		}

		clientsPic = append(clientsPic, clientPic)
	}

	svc.rpo.CreateClientPic(ctx, tx, &clientsPic)

	return model.ClientResponse{
		Id:      client.Id,
		Name:    client.Name,
		Address: client.Address,
		Phone:   client.Phone,
	}
}

func (svc *Service) GetAllClients(ctx context.Context) []model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	clients := svc.rpo.GetAllClients(ctx, tx)

	var result []model.ClientResponse
	if len(*clients) > 0 {
		for _, value := range *clients {
			client := model.ClientResponse{
				Id:      value.Id,
				Name:    value.Name,
				Address: value.Address,
				Phone:   value.Phone,
			}

			result = append(result, client)
		}
	}

	return result
}
