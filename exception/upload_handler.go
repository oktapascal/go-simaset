package exception

import (
	"encoding/json"
	"github.com/oktapascal/go-simpro/config"
	"github.com/oktapascal/go-simpro/web"
	"net/http"
)

type UploadFileError struct {
	Error string `json:"error"`
}

func NewUploadFileError(error string) UploadFileError {
	return UploadFileError{Error: error}
}

func UploadHandler(writer http.ResponseWriter, err any) {
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
