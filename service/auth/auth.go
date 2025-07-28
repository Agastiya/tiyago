package service

import (
	"errors"
	"fmt"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthService interface {
	LoginByEmail(params dto.LoginByEmailRequest) response.RespResultService
	RefreshToken(params dto.RefreshTokenRequest) response.RespResultService
}

func (s *AuthService) LoginByEmail(params dto.LoginByEmailRequest) response.RespResultService {

	user, err := s.UserRepo.DetailUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ResponseService(true, errors.New("email or password incorrect"), constant.StatusUnauthorized, nil, nil)
		}
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	if !user.Active {
		return response.ResponseService(true, errors.New("account disabled"), constant.StatusForbidden, nil, nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return response.ResponseService(true, errors.New("email or password incorrect"), constant.StatusUnauthorized, nil, nil)
	}

	loginResponse := dto.LoginResponse{
		Id:       user.Id,
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
	}

	accessToken, err := s.Jwt.GenerateToken(loginResponse, "secret_key")
	if err != nil {
		errMsg := fmt.Sprintf("failed to generate token. error : %v", err)
		return response.ResponseService(true, errors.New(errMsg), constant.StatusInternalServerError, nil, nil)
	}

	refreshToken, err := s.Jwt.GenerateToken(loginResponse, "refresh_secret_key")
	if err != nil {
		errMsg := fmt.Sprintf("failed to generate token. error : %v", err)
		return response.ResponseService(true, errors.New(errMsg), constant.StatusInternalServerError, nil, nil)
	}

	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, loginResponse)
}

func (s *AuthService) RefreshToken(params dto.RefreshTokenRequest) response.RespResultService {

	claims, err := s.Jwt.VerifyToken(params.RefreshToken, "refresh_secret_key")
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	claimsData := utils.MapClaimsToContextMap(claims)
	loginResponse := dto.LoginResponse{
		Id:       utils.StringToInt64(claimsData.Id),
		Fullname: claimsData.Fullname,
		Username: claimsData.Username,
		Email:    claimsData.Email,
	}

	accessToken, err := s.Jwt.GenerateToken(loginResponse, "secret_key")
	if err != nil {
		errMsg := fmt.Sprintf("failed to generate token. error : %v", err)
		return response.ResponseService(true, errors.New(errMsg), constant.StatusInternalServerError, nil, nil)
	}

	refreshToken, err := s.Jwt.GenerateToken(loginResponse, "refresh_secret_key")
	if err != nil {
		errMsg := fmt.Sprintf("failed to generate token. error : %v", err)
		return response.ResponseService(true, errors.New(errMsg), constant.StatusInternalServerError, nil, nil)
	}

	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, loginResponse)
}
