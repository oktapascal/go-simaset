package model

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	Client struct {
		Id           string
		Name         string
		Address      string
		Phone        string
		NumberOfPics int8
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

	UpdateClientPicRequest struct {
		Id      string `json:"id"`
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

	UpdateClientRequest struct {
		Id        string                   `json:"id" validate:"required"`
		Name      string                   `json:"name" validate:"required,min=1,max=50"`
		Address   string                   `json:"address" validate:"required,min=1,max=100"`
		Phone     string                   `json:"phone" validate:"required,min=11,max=13"`
		ClientPic []UpdateClientPicRequest `json:"client_pic" validate:"required,minclientpic,dive"`
	}

	ClientResponse struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Address      string `json:"address"`
		Phone        string `json:"phone"`
		NumberOfPics int8   `json:"number_of_pics"`
	}

	ClientPicResponse struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Address string `json:"address"`
	}

	ClientDetailResponse struct {
		Id        string              `json:"id"`
		Name      string              `json:"name"`
		Address   string              `json:"address"`
		Phone     string              `json:"phone"`
		ClientPic []ClientPicResponse `json:"client_pic"`
	}

	ClientRepository interface {
		GenerateClientKode(ctx context.Context, tx *sql.Tx) *string
		CreateClient(ctx context.Context, tx *sql.Tx, data *Client) *Client
		CreateClientPic(ctx context.Context, tx *sql.Tx, data *[]ClientPic) *[]ClientPic
		UpdateClient(ctx context.Context, tx *sql.Tx, data *Client) *Client
		UpdateClientPic(ctx context.Context, tx *sql.Tx, data *[]ClientPic) *[]ClientPic
		GetAllClients(ctx context.Context, tx *sql.Tx) *[]Client
		GetClient(ctx context.Context, tx *sql.Tx, id string) (*Client, error)
		GetClientPic(ctx context.Context, tx *sql.Tx, id string) *[]ClientPic
		DeleteClientPic(ctx context.Context, tx *sql.Tx, id string, clientId []string)
		DeleteClient(ctx context.Context, tx *sql.Tx, id string)
	}

	ClientService interface {
		StoreClient(ctx context.Context, request *SaveClientRequest) ClientResponse
		UpdateClient(ctx context.Context, request *UpdateClientRequest) ClientResponse
		GetAllClients(ctx context.Context) []ClientResponse
		GetOneClient(ctx context.Context, id string) ClientDetailResponse
		DeleteClient(ctx context.Context, id string)
	}

	ClientHandler interface {
		SaveClient() http.HandlerFunc
		UpdateClient() http.HandlerFunc
		GetAllClients() http.HandlerFunc
		GetOneClient() http.HandlerFunc
		DeleteClient() http.HandlerFunc
	}
)
