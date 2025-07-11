package routes

import (
	"github.com/agastiya/tiyago/controller"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Routes struct {
	Env        string
	Controller controller.Controller
	// Controller controller.ControllerInterface
	// Middleware Middleware.MiddlewareInterface
}

func (app *Routes) InitRoutes() *chi.Mux {

	appRoute := chi.NewRouter()
	appRoute.Use(chiMiddleware.RequestID)
	appRoute.Use(chiMiddleware.RealIP)
	appRoute.Use(chiMiddleware.RedirectSlashes)
	appRoute.Use(chiMiddleware.Recoverer)

	appRoute.Route("/tiyago", func(appRoute chi.Router) {
		appRoute.Group(func(appRoute chi.Router) {
			appRoute.Route("/user", func(appRoute chi.Router) {
				appRoute.Post("/", app.Controller.UserController.UserCreate)
			})
		})
		appRoute.Get("/ping", app.Controller.BaseController.Ping)
	})

	return appRoute
}
