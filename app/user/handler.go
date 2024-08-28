package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simaset/helper"
	"github.com/oktapascal/go-simaset/model"
	"github.com/oktapascal/go-simaset/web"
	"net/http"
)

type Handler struct {
	svc      model.UserService
	validate *validator.Validate
}

func (hdl *Handler) SaveUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.SaveUserRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		user := hdl.svc.SaveUser(ctx, req)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   user,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}
