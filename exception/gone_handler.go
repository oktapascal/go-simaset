package exception

import (
	"encoding/json"
	"github.com/oktapascal/go-simaset/config"
	"github.com/oktapascal/go-simaset/web"
	"net/http"
)

type GoneError struct {
	Error string
}

func NewGoneError(error string) GoneError {
	return GoneError{Error: error}
}

func GoneHandler(writer http.ResponseWriter, err any) {
	log := config.CreateLoggers(nil)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusGone)

	errorResponse := web.ErrorResponse{
		Code:   http.StatusGone,
		Status: http.StatusText(http.StatusGone),
		Errors: err,
	}

	encoder := json.NewEncoder(writer)

	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		log.Error(errEncoder)
	}
}
