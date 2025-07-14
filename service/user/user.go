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
	UpdateUser(params dto.UpdateUserRequest) response.RespResultService
	DeleteUser(params dto.DeleteUserRequest) response.RespResultService
}

func (s *UserService) CreateUser(params dto.CreateUserRequest) response.RespResultService {

	if err := utils.CheckExists("username", params.Username, 0, s.UserRepo.CheckUsernameExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	if err := utils.CheckExists("email", params.Email, 0, s.UserRepo.CheckEmailExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return response.ResponseService(true, errors.New("hashed password failed"), constant.StatusDataBadRequest, nil, nil)
	}

	userModel := &models.User{
		Fullname:  params.Username,
		Username:  params.Username,
		Email:     params.Email,
		Password:  string(hashedPassword),
		Active:    true,
		CreatedBy: "System",
		CreatedAt: utils.TimeNow(),
	}

	err = s.UserRepo.CreateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}

func (s *UserService) UpdateUser(params dto.UpdateUserRequest) response.RespResultService {

	if err := utils.CheckExists("username", params.Username, params.Id, s.UserRepo.CheckUsernameExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	if err := utils.CheckExists("email", params.Email, params.Id, s.UserRepo.CheckEmailExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	modifiedBy := "System"
	time := utils.TimeNow()
	userModel := &models.User{
		Id:         params.Id,
		Fullname:   params.Fullname,
		Username:   params.Username,
		Email:      params.Email,
		ModifiedBy: &modifiedBy,
		ModifiedAt: &time,
	}

	err := s.UserRepo.UpdateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}

func (s *UserService) DeleteUser(params dto.DeleteUserRequest) response.RespResultService {

	deletedBy := "System"
	time := utils.TimeNow()
	userModel := &models.User{
		Id:        params.Id,
		DeletedBy: &deletedBy,
		DeletedAt: &time,
	}

	err := s.UserRepo.DeleteUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}
