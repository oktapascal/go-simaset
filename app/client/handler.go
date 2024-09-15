package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"net/http"
)

type Handler struct {
	svc      model.ClientService
	validate *validator.Validate
}

func (hdl *Handler) SaveClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.SaveClientRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.RegisterValidation("minclientpic", func(fl validator.FieldLevel) bool {
			fmt.Println(fl.Field().String())
			return len(fl.Field().Interface().([]model.SaveClientPicRequest)) >= 1
		})
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		fmt.Println(req)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}
