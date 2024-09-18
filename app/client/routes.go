package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ClientHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/client", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)
		route.Get("/all", router.hdl.GetAllClients())
		route.Get("/{id}", router.hdl.GetOneClient())
		route.Post("/", router.hdl.SaveClient())
		route.Put("/{id}", router.hdl.UpdateClient())
		route.Delete("/{id}", router.hdl.DeleteClient())
	})
}
