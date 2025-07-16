package service

import (
	"github.com/agastiya/tiyago/pkg/jwt"
	"github.com/agastiya/tiyago/repository/user"
)

type AuthService struct {
	UserRepo user.IUserRepository
	Jwt      jwt.IJwt
}

type AuthServiceDeps struct {
	UserRepo user.IUserRepository
	Jwt      jwt.IJwt
}

func NewAuthService(deps AuthServiceDeps) IAuthService {
	return &AuthService{
		UserRepo: deps.UserRepo,
		Jwt:      deps.Jwt,
	}
}
