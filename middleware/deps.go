package middleware

import (
	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/jwt"
)

type Middleware struct {
	Jwt     jwt.IJwt
	Swagger dto.SwaggerSetting
}

type MiddlewareDeps struct {
	Jwt     jwt.IJwt
	Swagger dto.SwaggerSetting
}

func NewMiddleware(m MiddlewareDeps) IMiddleware {
	return &Middleware{
		Jwt:     m.Jwt,
		Swagger: m.Swagger,
	}
}
