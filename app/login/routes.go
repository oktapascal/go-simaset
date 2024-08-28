package login

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simaset/middleware"
	"github.com/oktapascal/go-simaset/model"
)

type Router struct {
	hdl model.LoginSessionHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/auth", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)

		route.Post("/login", router.hdl.Login())
		route.Post("/logout", router.hdl.Logout())
	})
}
