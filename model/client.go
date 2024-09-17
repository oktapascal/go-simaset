package model

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	Client struct {
		Id      string
		Name    string
		Address string
		Phone   string
	}

	ClientPic struct {
		Id       string
		ClientId string
		Name     string
		Phone    string
		Email    string
		Address  string
	}

	SaveClientPicRequest struct {
		Name    string `json:"name" validate:"required,min=1,max=50"`
		Email   string `json:"email" validate:"required,email,min=1,max=50"`
		Phone   string `json:"phone" validate:"required,min=1,max=13"`
		Address string `json:"address" validate:"required,min=1,max=100"`
	}

	SaveClientRequest struct {
		Name      string                 `json:"name" validate:"required,min=1,max=50"`
		Address   string                 `json:"address" validate:"required,min=1,max=100"`
		Phone     string                 `json:"phone" validate:"required,min=11,max=13"`
		ClientPic []SaveClientPicRequest `json:"client_pic" validate:"required,minclientpic,dive"`
	}

	ClientResponse struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	ClientRepository interface {
		CreateClient(ctx context.Context, tx *sql.Tx, data *Client) *Client
		CreateClientPic(ctx context.Context, tx *sql.Tx, data *[]ClientPic) *[]ClientPic
		GetAllClients(ctx context.Context, tx *sql.Tx) *[]Client
	}

	ClientService interface {
		StoreClient(ctx context.Context, request *SaveClientRequest) ClientResponse
		GetAllClients(ctx context.Context) []ClientResponse
	}

	ClientHandler interface {
		SaveClient() http.HandlerFunc
		GetAllClients() http.HandlerFunc
	}
)
