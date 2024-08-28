package helper

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/schema"
	"net/http"
)

var schemaDecoder = schema.NewDecoder()

func DecodeRequest(request *http.Request, out any) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&out)
	if err != nil && !errors.Is(err, http.ErrBodyReadAfterClose) {
		return err
	}

	err = request.ParseForm()
	if err != nil {
		return err
	}

	err = schemaDecoder.Decode(out, request.Form)
	if err != nil {
		return err
	}

	return nil
}
