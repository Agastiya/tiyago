package service

import (
	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/models"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
	user "github.com/agastiya/tiyago/repository/user"
)

type (
	UserService struct {
		Repo user.IUserRepository
	}

	IUserService interface {
		CreateUser(params dto.CreateUserRequest) response.RespResultService
	}
)

func NewUserService(repo user.IUserRepository) IUserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(params dto.CreateUserRequest) response.RespResultService {

	userModel := &models.User{
		Fullname:  params.Fullname,
		Email:     params.Email,
		Password:  "hashed_password",
		Active:    false,
		CreatedBy: "system",
		CreatedAt: utils.TimeNow(),
	}

	_, err := s.Repo.CreateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}
