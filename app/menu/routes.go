package menu

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.MenuHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/menu", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)
		route.Get("/all", router.hdl.GetMenu())
	})
}
