package middleware

import (
	"github.com/agastiya/tiyago/pkg/jwt"
)

type Middleware struct {
	Jwt jwt.IJwt
}

func NewMiddleware(jwt jwt.IJwt) IMiddleware {
	return &Middleware{Jwt: jwt}
}
