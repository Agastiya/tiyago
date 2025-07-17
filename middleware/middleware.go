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
	UserAuth() func(http.Handler) http.Handler
}

func (m Middleware) UserAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				response.ResponseError(w, errors.New("token not exists"), constant.StatusUnauthorized)
				return
			}

			if !strings.Contains(tokenString, "Bearer ") {
				response.ResponseError(w, errors.New("bearer not found"), constant.StatusUnauthorized)
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			claims, err := m.Jwt.VerifyToken(tokenString, "secret_key")
			if err != nil {
				errMessage := fmt.Sprintf("error verifying token: %v", err)
				response.ResponseError(w, errors.New(errMessage), constant.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), constant.ClaimsKey, claims)
			r.Header.Set("claims_value", fmt.Sprintf("%v", claims))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
