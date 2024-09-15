package model

import "net/http"

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

	saveClientPicRequest struct {
		Name    string `json:"name" validate:"required,min=1,max=50"`
		Email   string `json:"email" validate:"required,email,min=1,max=50"`
		Phone   string `json:"phone" validate:"required,min=1,max=13"`
		Address string `json:"address" validate:"required,min=1,max=100"`
	}

	SaveClientRequest struct {
		Name      string                 `json:"name" validate:"required,min=1,max=50"`
		Address   string                 `json:"address" validate:"required,min=1,max=100"`
		Phone     string                 `json:"phone" validate:"required,min=1,max=13"`
		ClientPic []saveClientPicRequest `json:"client_pic" validate:"required,dive"`
	}

	ClientResponse struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	ClientRepository interface{}

	ClientService interface{}

	ClientHandler interface {
		SaveClient() http.HandlerFunc
	}
)
