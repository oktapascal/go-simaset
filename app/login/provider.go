package login

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/oktapascal/go-simpro/app/user"
	"github.com/oktapascal/go-simpro/model"
	"sync"
)

var (
	route     *Router
	routeOnce sync.Once

	hdl     *Handler
	hdlOnce sync.Once

	svc     *Service
	svcOnce sync.Once

	rpo     *Repository
	rpoOnce sync.Once

	ProviderSet = wire.NewSet(
		ProvideRoute,
		ProvideHandler,
		ProvideService,
		ProvideRepository,
		user.ProvideRepository,
		wire.Bind(new(model.LoginSessionHandler), new(*Handler)),
		wire.Bind(new(model.LoginSessionService), new(*Service)),
		wire.Bind(new(model.LoginSessionRepository), new(*Repository)),
		wire.Bind(new(model.UserRepository), new(*user.Repository)),
	)
)

func ProvideRoute(hdl model.LoginSessionHandler) *Router {
	routeOnce.Do(func() {
		route = &Router{
			hdl: hdl,
		}
	})

	return route
}

func ProvideHandler(svc model.LoginSessionService, validate *validator.Validate) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc:      svc,
			validate: validate,
		}
	})

	return hdl
}

func ProvideService(rpo model.LoginSessionRepository, urpo model.UserRepository, db *sql.DB) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			rpo:  rpo,
			urpo: urpo,
			db:   db,
		}
	})

	return svc
}

func ProvideRepository() *Repository {
	rpoOnce.Do(func() {
		rpo = new(Repository)
	})

	return rpo
}
