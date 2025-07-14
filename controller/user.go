package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
	userSvc "github.com/agastiya/tiyago/service/user"
	"github.com/gorilla/schema"
)

type UserController struct {
	UserService userSvc.IUserService
}

func NewUserController(service userSvc.IUserService) IUserController {
	return &UserController{UserService: service}
}

type IUserController interface {
	UserBrowse(w http.ResponseWriter, r *http.Request)
	UserDetail(w http.ResponseWriter, r *http.Request)
	UserCreate(w http.ResponseWriter, r *http.Request)
	UserUpdate(w http.ResponseWriter, r *http.Request)
	UserDelete(w http.ResponseWriter, r *http.Request)
}

func (uc *UserController) UserBrowse(w http.ResponseWriter, r *http.Request) {

	var params dto.BrowseUserRequest
	err := schema.NewDecoder().Decode(&params, r.URL.Query())
	if err != nil {
		response.ResponseError(w, errors.New("parameter tidak valid"), constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.BrowseUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}

func (uc *UserController) UserDetail(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetUrl(r, "id")
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.DetailUser(id)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
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

func (uc *UserController) UserUpdate(w http.ResponseWriter, r *http.Request) {

	var params dto.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	params.Id, err = utils.GetUrl(r, "id")
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.UpdateUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}

func (uc *UserController) UserDelete(w http.ResponseWriter, r *http.Request) {

	var params dto.DeleteUserRequest
	var err error

	params.Id, err = utils.GetUrl(r, "id")
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.DeleteUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseError(w, nil, constant.StatusOKJson)
}
