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
	UserUpdatePassword(w http.ResponseWriter, r *http.Request)
	UserDelete(w http.ResponseWriter, r *http.Request)
}

// @Tags		User
// @Summary		Browse Users
// @Description Sample Parameter: `?sortColumn=id&sortOrder=desc&pageSize=20`
// @Accept		json
// @Produce		json
// @Param		sortColumn		query	string		false	"sortColumn"
// @Param		sortOrder		query	string		false	"sortOrder"
// @Param		pageSize		query	int64		false	"pageSize"
// @Param		pageNumber		query	int64		false	"pageNumber"
// @Param		fullname		query	string		false	"fullname"
// @Param		username		query	string		false	"username"
// @Param		email			query	string		false	"email"
// @Security	Bearer
// @Router		/user [get]
func (uc *UserController) UserBrowse(w http.ResponseWriter, r *http.Request) {
	var params dto.BrowseUserRequest
	err := schema.NewDecoder().Decode(&params, r.URL.Query())
	if err != nil {
		response.ResponseError(w, errors.New("invalid parameter"), constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.BrowseUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}

// @Tags		User
// @Summary		Detail user
// @Description Sample Parameter: `1`
// @Accept		json
// @Produce		json
// @Param		id	path	int64	true	"id"
// @Security	Bearer
// @Router		/user/{id} [get]
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

// @Tags        User
// @Summary     Create User
// @Description Example value: `{"fullname":"Agastiya","username":"ageztya777","email":"ageztya.putra@gmail.com","password":"12345678"}`
// @Accept      json
// @Produce     json
// @Param       "request body"	body	dto.CreateUserRequest	true "example payload"
// @Security	Bearer
// @Router    	/user [post]
func (uc *UserController) UserCreate(w http.ResponseWriter, r *http.Request) {
	var params dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	if err := utils.Validate(params); err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	ctxData, _, err := utils.GetUserClaimsFromContext(r)
	if err != nil {
		response.ResponseError(w, err, constant.StatusInternalServerError)
		return
	}

	params.CreatedBy = ctxData.Fullname
	result := uc.UserService.CreateUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}

// @Tags        User
// @Summary     Update User
// @Description Example value: `id:2, body:{"email":"jhon.doe@gmail.com","fullname":"Jhon Doe","username":"jhon.doe"}`
// @Accept      json
// @Produce     json
// @Param		id					path	int64					true	"id"
// @Param       "request 	body"	body	dto.UpdateUserRequest	true 	"example payload"
// @Security	Bearer
// @Router    	/user/{id} [put]
func (uc *UserController) UserUpdate(w http.ResponseWriter, r *http.Request) {

	var params dto.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	if err := utils.Validate(params); err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	ctxData, _, err := utils.GetUserClaimsFromContext(r)
	if err != nil {
		response.ResponseError(w, err, constant.StatusInternalServerError)
		return
	}

	params.Id, err = utils.GetUrl(r, "id")
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	params.ModifiedBy = ctxData.Fullname
	result := uc.UserService.UpdateUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}

// @Tags        User
// @Summary     Update User Password
// @Description Example value: `body:{"oldPassword":"1234567890","newPassword":"Aa123456!"}`
// @Accept      json
// @Produce     json
// @Param       "request 	body"	body	dto.UpdateUserPasswordRequest	true 	"example payload"
// @Security	Bearer
// @Router    	/user/password [put]
func (uc *UserController) UserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var params dto.UpdateUserPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.ResponseError(w, errors.New("invalid parameter"), constant.StatusDataBadRequest)
		return
	}

	ctxData, _, err := utils.GetUserClaimsFromContext(r)
	if err != nil {
		response.ResponseError(w, err, constant.StatusInternalServerError)
		return
	}
	params.UserId = utils.StringToInt64(ctxData.Id)

	if err := utils.Validate(params); err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	result := uc.UserService.UpdateUserPassword(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)

}

// @Tags        User
// @Summary     Delete User
// @Description Example value: `1`
// @Accept      json
// @Produce     json
// @Param		id					path	int64					true	"id"
// @Security	Bearer
// @Router    	/user/{id} [delete]
func (uc *UserController) UserDelete(w http.ResponseWriter, r *http.Request) {

	var params dto.DeleteUserRequest
	var err error

	ctxData, _, err := utils.GetUserClaimsFromContext(r)
	if err != nil {
		response.ResponseError(w, err, constant.StatusInternalServerError)
		return
	}

	params.Id, err = utils.GetUrl(r, "id")
	if err != nil {
		response.ResponseError(w, err, constant.StatusDataBadRequest)
		return
	}

	params.DeletedBy = ctxData.Fullname
	result := uc.UserService.DeleteUser(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseError(w, nil, constant.StatusOKJson)
}
