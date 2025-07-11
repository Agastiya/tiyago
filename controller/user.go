package controller

import (
	"encoding/json"
	"net/http"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	userSvc "github.com/agastiya/tiyago/service/user"
)

type UserController struct {
	UserService userSvc.IUserService
}

func NewUserController(service userSvc.IUserService) IUserController {
	return &UserController{UserService: service}
}

type IUserController interface {
	UserCreate(w http.ResponseWriter, r *http.Request)
}

func (uc *UserController) UserCreate(w http.ResponseWriter, r *http.Request) {

	var params dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.CreateUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}
