package service

import (
	"errors"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/models"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(params dto.CreateUserRequest) response.RespResultService
}

func (s *UserService) CreateUser(params dto.CreateUserRequest) response.RespResultService {

	if params.Username != nil {
		if err := utils.CheckExists("username", *params.Username, s.UserRepo.CheckUsernameExists); err != nil {
			return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
		}
	}

	if err := utils.CheckExists("email", params.Email, s.UserRepo.CheckEmailExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return response.ResponseService(true, errors.New("hashed password failed"), constant.StatusDataBadRequest, nil, nil)
	}

	userModel := &models.User{
		Fullname:  params.Fullname,
		Username:  params.Username,
		Email:     params.Email,
		Password:  string(hashedPassword),
		Active:    true,
		CreatedBy: "System",
		CreatedAt: utils.TimeNow(),
	}

	_, err = s.UserRepo.CreateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}
