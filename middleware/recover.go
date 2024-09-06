package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/exception"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if str, ok := err.(exception.NotFoundError); ok {
					exception.NotFoundHandler(writer, str)
					return
				}

				if str, ok := err.(exception.DuplicateError); ok {
					exception.DuplicateHandler(writer, str)
					return
				}

				if str, ok := err.(exception.NotMatchedError); ok {
					exception.NotMatchedHandler(writer, str)
					return
				}

				if str, ok := err.(exception.GoneError); ok {
					exception.GoneHandler(writer, str)
					return
				}

				if str, ok := err.(exception.PermissionError); ok {
					exception.PermissionHandler(writer, str)
					return
				}

				if str, ok := err.(exception.ForbiddenError); ok {
					exception.ForbiddenHandler(writer, str)
					return
				}

				if str, ok := err.(string); ok {
					exception.InternalServerHandler(writer, str)
					return
				}

				var validationErrors validator.ValidationErrors
				if errors.As(err.(error), &validationErrors) {
					badRequestHandler := exception.FormatErrors(validationErrors)
					exception.BadRequestHandler(writer, badRequestHandler)
					return
				}

				exception.InternalServerHandler(writer, err)
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
