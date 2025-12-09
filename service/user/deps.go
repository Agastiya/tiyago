package service

import "github.com/agastiya/tiyago/repository/user"

type UserService struct {
	UserRepo user.IUserRepository
}

type UserServiceDeps struct {
	UserRepo user.IUserRepository
}

func NewUserService(deps UserServiceDeps) IUserService {
	return &UserService{
		UserRepo: deps.UserRepo,
	}
}
