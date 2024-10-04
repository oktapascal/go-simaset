package menu

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
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
		wire.Bind(new(model.MenuHandler), new(*Handler)),
		wire.Bind(new(model.MenuService), new(*Service)),
		wire.Bind(new(model.MenuRepository), new(*Repository)),
	)
)

func ProvideRoute(hdl model.MenuHandler) *Router {
	routeOnce.Do(func() {
		route = &Router{
			hdl: hdl,
		}
	})

	return route
}

func ProvideHandler(svc model.MenuService, validate *validator.Validate) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc:      svc,
			validate: validate,
		}
	})

	return hdl
}

func ProvideService(rpo model.MenuRepository) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			rpo: rpo,
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
