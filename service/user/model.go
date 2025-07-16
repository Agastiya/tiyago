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

var allowedFieldToSort = map[string]string{
	"id":         "id",
	"fullname":   "fullname",
	"username":   "username",
	"email":      "email",
	"createdBy":  "created_by",
	"createdAt":  "created_at",
	"modifiedBy": "modified_by",
	"modifiedAt": "modified_at",
}
