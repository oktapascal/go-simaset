package middleware

import (
	"net/http"
	"strings"
)

func AuthorizationCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if !strings.Contains(authorization, "Bearer") {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		next.ServeHTTP(w, r)
	})
}
