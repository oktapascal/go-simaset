package menu_group

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.MenuGroupHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/menu-group", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)
		route.Get("/all", router.hdl.GetAllMenuGroups())
		route.Get("/{id}", router.hdl.GetOneMenuGroup())
		route.Post("/", router.hdl.SaveMenuGroup())
		route.Put("/{id}", router.hdl.UpdateMenuGroup())
		route.Delete("/{id}", router.hdl.DeleteMenuGroup())
	})
}
