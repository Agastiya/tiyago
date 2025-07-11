package service

import "github.com/agastiya/tiyago/repository/user"

type UserService struct {
	UserRepo user.IUserRepository
	AuthRepo user.IUserRepository
}

type UserServiceDeps struct {
	UserRepo user.IUserRepository
	AuthRepo user.IUserRepository
}

func NewUserService(deps UserServiceDeps) IUserService {
	return &UserService{
		UserRepo: deps.UserRepo,
		AuthRepo: deps.AuthRepo,
	}
}
