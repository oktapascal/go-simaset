package exception

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simaset/web"
	"net/http"
)

type FormatError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func convertTagToMessage(ex validator.FieldError) string {
	switch ex.Tag() {
	case "required":
		return "kolom ini tidak boleh kosong"
	case "email":
		return "email tidak valid"
	case "min":
		return fmt.Sprint("kolom ini harus memiliki panjang minimal ", ex.Param(), " karakter")
	case "max":
		return fmt.Sprint("kolom ini harus memiliki panjang maksimal ", ex.Param(), " karakter")
	case "eqfield":
		return fmt.Sprint("kolom ini harus sama dengan '", ex.Param(), "'")
	case "oneof":
		return fmt.Sprintf("kolom ini harus terisi dari: %s", ex.Param())
	case "url":
		return "URL tidak valid"
	default:
		return ex.Error()
	}
}

func FormatErrors(error error) []FormatError {
	var exception validator.ValidationErrors

	if errors.As(error, &exception) {
		fieldErrors := make([]FormatError, len(exception))

		for index, ex := range exception {
			fieldErrors[index] = FormatError{
				Param:   ex.Field(),
				Message: convertTagToMessage(ex),
			}
		}

		return fieldErrors
	}

	return nil
}

func BadRequestHandler(writer http.ResponseWriter, error any) {
	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusBadRequest)

	responseError := web.ErrorResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: error,
	}

	encoder := json.NewEncoder(writer)

	if err := encoder.Encode(responseError); err != nil {
		InternalServerHandler(writer, err)
	}
}
