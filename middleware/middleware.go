package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
)

type IMiddleware interface {
	JWTAuth() func(http.Handler) http.Handler
	SwaggerAuth() func(http.Handler) http.Handler
}

func (m Middleware) JWTAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				response.JSONResponse(w, nil, errors.New("token not exists"), http.StatusUnauthorized)
				return
			}

			if !strings.Contains(tokenString, "Bearer ") {
				response.JSONResponse(w, nil, errors.New("bearer not found"), http.StatusUnauthorized)
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			claims, err := m.Jwt.VerifyToken(tokenString, "secret_key")
			if err != nil {
				errMessage := fmt.Sprintf("error: %v", err)
				response.JSONResponse(w, nil, errors.New(errMessage), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), constant.ClaimsKey, claims)
			r.Header.Set("claims_value", fmt.Sprintf("%v", claims))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (m Middleware) SwaggerAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if !ok || username != m.Swagger.Username || password != m.Swagger.Password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
