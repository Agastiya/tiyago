package service

import (
	"errors"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/models"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/helper/response"
	"github.com/agastiya/tiyago/pkg/helper/utils"
	"github.com/agastiya/tiyago/repository/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	BrowseUser(params dto.BrowseUserRequest) response.RespResultService
	DetailUser(id int64) response.RespResultService
	CreateUser(params dto.CreateUserRequest) response.RespResultService
	UpdateUser(params dto.UpdateUserRequest) response.RespResultService
	UpdateUserPassword(params dto.UpdateUserPasswordRequest) response.RespResultService
	DeleteUser(params dto.DeleteUserRequest) response.RespResultService
}

func (s *UserService) BrowseUser(params dto.BrowseUserRequest) response.RespResultService {

	var defaultPaginationParams = dto.Pagination{PageSize: 10, PageNumber: 0, SortColumn: "id", SortOrder: "DESC"}
	params.Pagination = utils.SetDefaultParams(params.Pagination, defaultPaginationParams)
	params.SortOrder = utils.ValidateSortOrder(params.SortOrder, defaultPaginationParams.SortOrder)
	params.SortColumn = utils.ValidateSortColumn(allowedFieldToSort, params.SortColumn, defaultPaginationParams.SortColumn)

	filter := user.BrowseUserFilter{
		PageSize:   params.PageSize,
		PageNumber: params.PageNumber * params.PageSize,
		SortColumn: params.SortColumn,
		SortOrder:  params.SortOrder,
		Fullname:   params.Fullname,
		Username:   params.Username,
		Email:      params.Email,
	}

	result, err := s.UserRepo.BrowseUser(filter)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	browseResult := make([]dto.UserResponse, len(result))
	var totalRecords int
	var hasReachMax bool

	for i, user := range result {
		browseResult[i] = dto.UserResponse{
			Id:         user.Id,
			Fullname:   user.Fullname,
			Username:   user.Username,
			Email:      user.Email,
			Active:     user.Active,
			CreatedBy:  user.CreatedBy,
			CreatedAt:  user.CreatedAt,
			ModifiedBy: user.ModifiedBy,
			ModifiedAt: user.ModifiedAt,
		}
		if i == 0 {
			totalRecords = user.TotalRecords
			hasReachMax = user.HasReachMax
		}
	}

	resultData := dto.BrowseModel[dto.UserResponse]{
		RecordsTotal: totalRecords,
		HasReachMax:  hasReachMax,
		Data:         browseResult,
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, resultData)
}

func (s *UserService) DetailUser(id int64) response.RespResultService {

	result, err := s.UserRepo.DetailUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ResponseService(true, err, constant.StatusDataNotFound, nil, nil)
		}
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	detailResult := dto.UserResponse{
		Id:         result.Id,
		Fullname:   result.Fullname,
		Username:   result.Username,
		Email:      result.Email,
		Active:     result.Active,
		CreatedBy:  result.CreatedBy,
		CreatedAt:  result.CreatedAt,
		ModifiedBy: result.ModifiedBy,
		ModifiedAt: result.ModifiedAt,
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, detailResult)
}

func (s *UserService) CreateUser(params dto.CreateUserRequest) response.RespResultService {

	if err := utils.CheckExistsFieldName("username", params.Username, 0, s.UserRepo.CheckUsernameExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	if err := utils.CheckExistsFieldName("email", params.Email, 0, s.UserRepo.CheckEmailExists); err != nil {
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
		CreatedBy: params.CreatedBy,
		CreatedAt: utils.TimeNow(),
	}

	err = s.UserRepo.CreateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}

func (s *UserService) UpdateUser(params dto.UpdateUserRequest) response.RespResultService {

	if err := utils.CheckExistsFieldName("username", params.Username, params.Id, s.UserRepo.CheckUsernameExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	if err := utils.CheckExistsFieldName("email", params.Email, params.Id, s.UserRepo.CheckEmailExists); err != nil {
		return response.ResponseService(true, err, constant.StatusDataBadRequest, nil, nil)
	}

	time := utils.TimeNow()
	userModel := &models.User{
		Id:         params.Id,
		Fullname:   params.Fullname,
		Username:   params.Username,
		Email:      params.Email,
		ModifiedBy: &params.ModifiedBy,
		ModifiedAt: &time,
	}

	err := s.UserRepo.UpdateUser(userModel)
	if err != nil {
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	return response.ResponseService(false, nil, constant.StatusOKJson, nil, nil)
}

func (s *UserService) UpdateUserPassword(params dto.UpdateUserPasswordRequest) response.RespResultService {

	user, err := s.UserRepo.DetailUser(params.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ResponseService(true, errors.New("user not found"), constant.StatusDataNotFound, nil, nil)
		}
		return response.ResponseService(true, err, constant.StatusInternalServerError, nil, nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.OldPassword))
	if err != nil {
		return response.ResponseService(true, errors.New("old password incorrect"), constant.StatusUnauthorized, nil, nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), 10)
	if err != nil {
		return response.ResponseService(true, errors.New("hashed new password failed"), constant.StatusDataBadRequest, nil, nil)
	}

	time := utils.TimeNow()
	userModel := &models.User{
		Id:         params.UserId,
		Password:   string(hashedPassword),
		ModifiedBy: &params.ModifiedBy,
		ModifiedAt: &time,
	}

	err = s.UserRepo.UpdateUserPassword(userModel)
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
