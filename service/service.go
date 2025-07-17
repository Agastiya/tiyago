package service

import (
	"github.com/agastiya/tiyago/pkg/jwt"
	"github.com/agastiya/tiyago/repository"
	authSvc "github.com/agastiya/tiyago/service/auth"
	userSvc "github.com/agastiya/tiyago/service/user"
)

type Service struct {
	User userSvc.IUserService
	Auth authSvc.IAuthService
}

type ServiceDeps struct {
	Repos   *repository.Repositories
	Package Package
}

type Package struct {
	Jwt jwt.IJwt
}

func InitServices(deps ServiceDeps) *Service {
	return &Service{
		User: userSvc.NewUserService(userSvc.UserServiceDeps{
			UserRepo: deps.Repos.UserRepo,
		}),
		Auth: authSvc.NewAuthService(authSvc.AuthServiceDeps{
			UserRepo: deps.Repos.UserRepo,
			Jwt:      deps.Package.Jwt,
		}),
	}
}
