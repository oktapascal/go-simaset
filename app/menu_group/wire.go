//go:build wireinject
// +build wireinject

package menu_group

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func Wire(validate *validator.Validate, db *sql.DB) *Router {
	panic(wire.Build(ProviderSet))
}
