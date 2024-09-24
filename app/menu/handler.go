package menu

import (
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/model"
	"net/http"
)

type Handler struct {
	svc      model.MenuService
	validate *validator.Validate
}

func (hdl *Handler) GetMenus() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (hdl *Handler) GetMenuChildren() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
