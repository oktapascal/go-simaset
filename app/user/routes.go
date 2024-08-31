package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simaset/middleware"
	"github.com/oktapascal/go-simaset/model"
)

type Router struct {
	hdl model.UserHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/user", func(route chi.Router) {
		route.Post("/", router.hdl.SaveUser())

		route.Group(func(route chi.Router) {
			route.Use(middleware.AuthorizationCheckMiddleware)
			route.Use(middleware.VerifyAccessTokenMiddleware)
			route.Get("/with-token", router.hdl.GetUserByToken())
		})
	})
}
