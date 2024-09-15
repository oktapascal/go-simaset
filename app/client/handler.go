package client

import (
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/model"
	"net/http"
)

type Handler struct {
	svc      model.ClientService
	validate *validator.Validate
}

func (hdl *Handler) SaveClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
