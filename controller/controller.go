package controller

import service "github.com/agastiya/tiyago/service"

type Controller struct {
	Base Base
	User IUserController
}

func NewController(service service.Service) Controller {
	return Controller{
		// Base: NewBaseController(),
		User: NewUserController(service.User),
	}
}
