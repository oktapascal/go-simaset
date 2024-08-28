package exception

import (
	"encoding/json"
	"github.com/oktapascal/go-simaset/config"
	"github.com/oktapascal/go-simaset/web"
	"net/http"
)

type DuplicateError struct {
	Error string
}

func NewDuplicateError(error string) DuplicateError {
	return DuplicateError{
		Error: error,
	}
}

func DuplicateHandler(writer http.ResponseWriter, err any) {
	log := config.CreateLoggers(nil)

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusBadRequest)

	errorResponse := web.ErrorResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: err,
	}

	encoder := json.NewEncoder(writer)

	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		log.Error(errEncoder)
	}
}
