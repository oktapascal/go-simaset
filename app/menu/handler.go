package menu

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"net/http"
)

type Handler struct {
	svc      model.MenuService
	validate *validator.Validate
}

func (hdl *Handler) GetMenu() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		group, ok := userInfo["menu_group"].(string)
		if !ok {
			panic("Something wrong when extracting menu group from jwt token")
		}

		result := hdl.svc.GetMenu(group)

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
