package login

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"net/http"
)

type Handler struct {
	svc      model.LoginSessionService
	validate *validator.Validate
}

func (hdl *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.LoginRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.Login(ctx, req, request.UserAgent())

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) Logout() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		ctx := request.Context()
		hdl.svc.Logout(ctx, userInfo, request.UserAgent())

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) GetAccessToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		ctx := request.Context()
		result := hdl.svc.GenerateAccessToken(ctx, userInfo, request.UserAgent())

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}
