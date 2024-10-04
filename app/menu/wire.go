//go:build wireinject
// +build wireinject

package menu

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func Wire(validate *validator.Validate) *Router {
	panic(wire.Build(ProviderSet))
}
