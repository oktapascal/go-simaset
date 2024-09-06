package login

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.LoginSessionHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/auth", func(route chi.Router) {
		route.Post("/login", router.hdl.Login())
		route.Group(func(route chi.Router) {
			route.Use(middleware.AuthorizationCheckMiddleware)
			route.Use(middleware.VerifyRefreshTokenMiddleware)
			route.Get("/access-token", router.hdl.GetAccessToken())
		})
		route.Group(func(route chi.Router) {
			route.Use(middleware.AuthorizationCheckMiddleware)
			route.Use(middleware.VerifyAccessTokenMiddleware)
			route.Post("/logout", router.hdl.Logout())
		})
	})
}
