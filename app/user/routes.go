package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simaset/model"
)

type Router struct {
	hdl model.UserHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/user", func(route chi.Router) {
		route.Post("/", router.hdl.SaveUser())
	})
}
