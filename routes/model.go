package routes

import (
	"github.com/agastiya/tiyago/controller"
	"github.com/agastiya/tiyago/middleware"
)

type Routes struct {
	Env        string
	Controller controller.Controller
	Middleware middleware.IMiddleware
}
