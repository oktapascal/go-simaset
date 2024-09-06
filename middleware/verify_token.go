package middleware

import (
	"context"
	"github.com/oktapascal/go-simpro/helper"
	"net/http"
	"strings"
)

func VerifyAccessTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")
		token := strings.Replace(authorization, "Bearer", "", -1)
		token = strings.TrimSpace(token)

		verify, err := helper.VerifyAccessToken(token)

		if err != nil {
			http.Error(writer, err.Error(), 401)
			return
		}

		claims := verify.Claims
		ctx := context.WithValue(request.Context(), "claims", claims)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func VerifyRefreshTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")
		token := strings.Replace(authorization, "Bearer", "", -1)
		token = strings.TrimSpace(token)

		verify, err := helper.VerifyRefreshToken(token)

		if err != nil {
			http.Error(writer, err.Error(), 401)
			return
		}

		claims := verify.Claims
		ctx := context.WithValue(request.Context(), "claims", claims)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
