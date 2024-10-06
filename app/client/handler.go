package client

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"net/http"
	"strconv"
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
			return len(fl.Field().Interface().([]model.SaveClientPicRequest)) >= 1
		})
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.StoreClient(ctx, req)

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

func (hdl *Handler) GetAllClients() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := helper.DefaultPaginationParams()

		pageParam := request.URL.Query().Get("page")
		pageSizeParam := request.URL.Query().Get("page_size")
		sortByParam := request.URL.Query().Get("sort_by")
		orderByParam := request.URL.Query().Get("order_by")
		filterByParam := request.URL.Query().Get("filter_by")
		filterValueParam := request.URL.Query().Get("filter_value")
		cursorParam := request.URL.Query().Get("cursor")

		page, errPageParam := strconv.Atoi(pageParam)
		if errPageParam != nil || page < 1 {
			page = params.Page
		}

		pageSize, errPageSize := strconv.Atoi(pageSizeParam)
		if errPageSize != nil || pageSize < 1 {
			pageSize = params.PageSize
		}

		params.ApplyPaginationParams(page, pageSize, sortByParam, orderByParam, filterByParam, filterValueParam, cursorParam)

		ctx := request.Context()
		result := hdl.svc.GetAllClients(ctx)

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

func (hdl *Handler) GetOneClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		ctx := request.Context()
		result := hdl.svc.GetOneClient(ctx, id)

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

func (hdl *Handler) UpdateClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.UpdateClientRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.RegisterValidation("minclientpic", func(fl validator.FieldLevel) bool {
			return len(fl.Field().Interface().([]model.UpdateClientPicRequest)) >= 1
		})
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.UpdateClient(ctx, req)

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

func (hdl *Handler) DeleteClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		ctx := request.Context()
		hdl.svc.DeleteClient(ctx, id)

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
