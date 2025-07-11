package controller

import service "github.com/agastiya/tiyago/service"

type Controller struct {
	BaseController IBaseController
	UserController IUserController
}

func InitController(service service.Service) *Controller {
	return &Controller{
		BaseController: NewBaseController(),
		UserController: NewUserController(service.User),
	}
}
