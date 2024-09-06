package middleware

import (
	"github.com/oktapascal/go-simpro/config"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log := config.CreateLoggers(request)

		log.Info("Incoming Request")

		next.ServeHTTP(writer, request)
	})
}
