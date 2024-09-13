package user

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
		result := hdl.svc.SaveUser(ctx, req)

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

func (hdl *Handler) GetUserByToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		ctx := request.Context()
		result := hdl.svc.GetUserByToken(ctx, userInfo)

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

func (hdl *Handler) EditUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)
		req := new(model.UpdateUserRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.EditUser(ctx, req, userInfo)

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

func (hdl *Handler) UploadPhotoProfile() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		ctx := request.Context()
		userId := hdl.svc.GetUserIdByToken(ctx, userInfo)

		const MaxSize = 2 * 1024 * 1024

		request.Body = http.MaxBytesReader(writer, request.Body, MaxSize)

		err := request.ParseMultipartForm(MaxSize)
		if err != nil {
			panic(exception.NewUploadFileError("file exceeds 2mb"))
		}

		file, header, errFile := request.FormFile("photo")
		if errFile != nil {
			panic(exception.NewUploadFileError("failed to retrieve file"))
		}

		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				panic(err.Error())
			}
		}(file)

		fileExt := strings.ToLower(filepath.Ext(header.Filename))

		if !(fileExt == ".png" || fileExt == ".jpg" || fileExt == ".jpeg") {
			panic(exception.NewUploadFileError("file format does not support image format"))
		}

		_, err = os.Stat("storage/applications/" + userId)
		if err != nil {
			if os.IsNotExist(err) {
				errMkdir := os.Mkdir("storage/applications/"+userId, os.ModePerm)
				if errMkdir != nil {
					log.Fatal(errMkdir)
				}
			}
		}

		fileName := "my-photo" + filepath.Ext(header.Filename)
		dst, errCreate := os.Create(filepath.Join("storage", "applications", userId, fileName))
		if errCreate != nil {
			panic(errCreate.Error())
		}

		defer func(dst *os.File) {
			err := dst.Close()
			if err != nil {
				panic(err.Error())
			}
		}(dst)

		_, errCopy := io.Copy(dst, file)
		if errCopy != nil {
			panic(errCopy.Error())
		}

		result := hdl.svc.EditPhotoUser(ctx, fileName, userInfo)

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

func (hdl *Handler) GetPhotoProfile() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username := chi.URLParam(request, "username")

		ctx := request.Context()
		user := hdl.svc.GetUserByUsername(ctx, username)

		path := filepath.Join("storage", "applications", user.Id, user.Photo)

		_, err := os.Stat(path)
		if err != nil {
			panic(exception.NewNotFoundError("file not found"))
		}

		http.ServeFile(writer, request, path)
	}
}
