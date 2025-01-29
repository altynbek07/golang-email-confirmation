package middleware

import (
	"context"
	"fmt"
	"go/email-verification/configs"
	"go/email-verification/pkg/jwt"
	"net/http"
	"strings"
)

type contextKey string

const (
	ContextEmailKey contextKey = "ContextEmailKey"
)

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware")
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(conf.Auth.Secret).Parse(token)
		if !isValid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
