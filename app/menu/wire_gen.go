// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package menu

import (
	"github.com/go-playground/validator/v10"
)

// Injectors from wire.go:

func Wire(validate *validator.Validate) *Router {
	repository := ProvideRepository()
	service := ProvideService(repository)
	handler := ProvideHandler(service, validate)
	router := ProvideRoute(handler)
	return router
}
