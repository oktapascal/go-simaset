package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.UserHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/user", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)
		route.Get("/with-token", router.hdl.GetUserByToken())
		route.Post("/", router.hdl.SaveUser())
		route.Put("/", router.hdl.EditUser())
		route.Post("/upload-photo", router.hdl.UploadPhotoProfile())
	})
}
