package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	authSvc "github.com/agastiya/tiyago/service/auth"
)

type AuthController struct {
	AuthService authSvc.IAuthService
}

func NewAuthController(service authSvc.IAuthService) IAuthController {
	return &AuthController{AuthService: service}
}

type IAuthController interface {
	LoginByEmail(w http.ResponseWriter, r *http.Request)
}

func (au *AuthController) LoginByEmail(w http.ResponseWriter, r *http.Request) {

	var params dto.LoginByEmailRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.ResponseError(w, errors.New("invalid parameter"), constant.StatusDataBadRequest)
		return
	}

	result := au.AuthService.LoginByEmail(params)
	if result.HasErr {
		response.ResponseError(w, result.Err, result.InternalCode)
		return
	}

	response.ResponseSuccess(w, result.Result, constant.StatusOKJson)
}
