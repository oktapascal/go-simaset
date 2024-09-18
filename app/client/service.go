package client

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/exception"
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

func (svc *Service) GetOneClient(ctx context.Context, id string) model.ClientDetailResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	client, errClient := svc.rpo.GetClient(ctx, tx, id)
	if errClient != nil {
		panic(exception.NewNotFoundError(errClient.Error()))
	}

	result := model.ClientDetailResponse{
		Id:        client.Id,
		Name:      client.Name,
		Address:   client.Address,
		Phone:     client.Phone,
		ClientPic: nil,
	}

	clientPic := svc.rpo.GetClientPic(ctx, tx, id)

	var clientsPic []model.ClientPicResponse
	if len(*clientPic) > 0 {
		for _, value := range *clientPic {
			cp := model.ClientPicResponse{
				Id:      value.Id,
				Name:    value.Name,
				Phone:   value.Phone,
				Email:   value.Email,
				Address: value.Address,
			}

			clientsPic = append(clientsPic, cp)
		}
	}

	result.ClientPic = clientsPic

	return result
}

func (svc *Service) UpdateClient(ctx context.Context, request *model.UpdateClientRequest) model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	client, errClient := svc.rpo.GetClient(ctx, tx, request.Id)
	if errClient != nil {
		panic(exception.NewNotFoundError(errClient.Error()))
	}

	client.Name = request.Name
	client.Phone = request.Phone
	client.Address = request.Address

	var clientPicCollections []string

	for _, value := range request.ClientPic {
		if value.Id != "" {
			clientPicCollections = append(clientPicCollections, value.Id)
		}
	}

	svc.rpo.DeleteClientPic(ctx, tx, request.Id, clientPicCollections)

	var clientsPicUpdate []model.ClientPic
	var clientsPicInsert []model.ClientPic

	for _, value := range request.ClientPic {
		if value.Id != "" {
			clientPicUpdate := model.ClientPic{
				Id:       value.Id,
				ClientId: client.Id,
				Name:     value.Name,
				Phone:    value.Phone,
				Email:    value.Email,
				Address:  value.Address,
			}

			clientsPicUpdate = append(clientsPicUpdate, clientPicUpdate)
		} else {
			clientPicInsert := model.ClientPic{
				ClientId: client.Id,
				Name:     value.Name,
				Phone:    value.Phone,
				Email:    value.Email,
				Address:  value.Address,
			}

			clientsPicInsert = append(clientsPicInsert, clientPicInsert)
		}
	}

	client = svc.rpo.UpdateClient(ctx, tx, client)
	svc.rpo.UpdateClientPic(ctx, tx, &clientsPicUpdate)

	if len(clientsPicInsert) > 0 {
		svc.rpo.CreateClientPic(ctx, tx, &clientsPicInsert)
	}

	return model.ClientResponse{
		Id:      client.Id,
		Name:    client.Name,
		Address: client.Address,
		Phone:   client.Phone,
	}
}

func (svc *Service) DeleteClient(ctx context.Context, id string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	svc.rpo.DeleteClient(ctx, tx, id)
}
