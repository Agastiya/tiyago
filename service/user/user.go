package service

import (
	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/models"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
)

type IUserService interface {
	CreateUser(params dto.CreateUserRequest) response.RespResultService
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

	_, err := s.UserRepo.CreateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}
