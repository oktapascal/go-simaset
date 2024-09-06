package welcome

import (
	"github.com/oktapascal/go-simpro/exception"
	"github.com/spf13/viper"
	"net/http"
)

type Handler struct {
}

// Welcome is a handler function that writes a welcome message to the HTTP response.
//
// It takes a http.ResponseWriter and a *http.Request as parameters.
// It does not return anything.
func (hdl *Handler) Welcome() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)

		_, err := writer.Write([]byte("Hello From " + viper.GetString("APP_NAME")))
		if err != nil {
			exception.InternalServerHandler(writer, err)
		}
	}
}

func (hdl *Handler) NotFoundApi() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)

		_, err := writer.Write([]byte("Route Doesn't Exist"))
		if err != nil {
			exception.InternalServerHandler(writer, err)
		}
	}
}

func (hdl *Handler) MethodNotAllowedApi() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusMethodNotAllowed)

		_, err := writer.Write([]byte("Method Is Not Allowed"))
		if err != nil {
			exception.InternalServerHandler(writer, err)
		}
	}
}
