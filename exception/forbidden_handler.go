package exception

import (
	"encoding/json"
	"github.com/oktapascal/go-simaset/config"
	"github.com/oktapascal/go-simaset/web"
	"net/http"
)

type ForbiddenError struct {
	Error string
}

func NewForbiddenError(error string) ForbiddenError {
	return ForbiddenError{Error: error}
}

func ForbiddenHandler(writer http.ResponseWriter, err any) {
	log := config.CreateLoggers(nil)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusForbidden)

	errorResponse := web.ErrorResponse{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
		Errors: err,
	}

	encoder := json.NewEncoder(writer)

	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		log.Error(errEncoder)
	}
}
