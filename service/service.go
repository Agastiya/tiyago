package service

import (
	"github.com/agastiya/tiyago/repository"
	userSvc "github.com/agastiya/tiyago/service/user"
)

type Service struct {
	User userSvc.IUserService
}

func InitServices(deps *repository.Repositories) *Service {
	return &Service{
		User: userSvc.NewUserService(userSvc.UserServiceDeps{
			UserRepo: deps.UserRepo,
		}),
	}
}
