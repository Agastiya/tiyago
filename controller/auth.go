package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
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
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

// @Tags        Auth
// @Summary     Login by Email & Password
// @Description Example value: `{"email":"tiyago@gmail.com","password":"tiyago12345"}`
// @Accept      json
// @Produce     json
// @Param       "request body" body    dto.LoginByEmailRequest   true  "Email & Password"
// @Router      /auth/loginbyemail [post]
func (au *AuthController) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	var params dto.LoginByEmailRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.JSONResponse(w, nil, errors.New("invalid parameter"), http.StatusBadRequest)
		return
	}

	if err := utils.Validate(params); err != nil {
		response.JSONResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	result := au.AuthService.LoginByEmail(params)
	if result.HasErr {
		response.JSONResponse(w, nil, result.Err, result.HttpCode)
		return
	}

	response.JSONResponse(w, result.Result, nil, result.HttpCode)
}

// @Tags        Auth
// @Summary     Refresh Token
// @Description Example value: `{"refreshToken":"qwerty1234567"}`
// @Accept      json
// @Produce     json
// @Param       "request body" body    dto.RefreshTokenRequest   true  "Refresh Token"
// @Router      /auth/refreshtoken [post]
func (au *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var params dto.RefreshTokenRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.JSONResponse(w, nil, errors.New("invalid parameter"), http.StatusBadRequest)
		return
	}

	if err := utils.Validate(params); err != nil {
		response.JSONResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	result := au.AuthService.RefreshToken(params)
	if result.HasErr {
		response.JSONResponse(w, nil, result.Err, result.HttpCode)
		return
	}

	response.JSONResponse(w, result.Result, nil, result.HttpCode)
}
