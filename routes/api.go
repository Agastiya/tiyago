package routes

import (
	"github.com/agastiya/tiyago/controller"
	_ "github.com/agastiya/tiyago/docs"
	"github.com/agastiya/tiyago/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	Env        string
	Controller controller.Controller
	Middleware middleware.IMiddleware
}

func (app *Routes) InitRoutes() *chi.Mux {

	appRoute := chi.NewRouter()
	appRoute.Use(chiMiddleware.RequestID)
	appRoute.Use(chiMiddleware.RealIP)
	appRoute.Use(chiMiddleware.RedirectSlashes)
	appRoute.Use(chiMiddleware.Recoverer)

	appRoute.Route("/tiyago", func(appRoute chi.Router) {
		appRoute.Route("/auth", func(appRoute chi.Router) {
			appRoute.Post("/loginbyemail", app.Controller.AuthController.LoginByEmail)
			appRoute.Post("/refreshtoken", app.Controller.AuthController.RefreshToken)
		})
		appRoute.Group(func(appRoute chi.Router) {
			appRoute.Use(app.Middleware.UserAuth())
			appRoute.Route("/user", func(appRoute chi.Router) {
				appRoute.Get("/", app.Controller.UserController.UserBrowse)
				appRoute.Post("/", app.Controller.UserController.UserCreate)
				appRoute.Route("/{id}", func(appRoute chi.Router) {
					appRoute.Get("/", app.Controller.UserController.UserDetail)
					appRoute.Put("/", app.Controller.UserController.UserUpdate)
					appRoute.Put("/password", app.Controller.UserController.UserUpdatePassword)
					appRoute.Delete("/", app.Controller.UserController.UserDelete)
				})
			})
		})

		appRoute.Get("/ping", app.Controller.BaseController.Ping)

		switch app.Env {
		case "local":
			appRoute.Mount("/swagger", httpSwagger.WrapHandler)
		case "development":
			appRoute.Group(func(appRoute chi.Router) {
				appRoute.With(app.Middleware.BasicAuthSwagger()).Mount("/swagger", httpSwagger.WrapHandler)
			})
		}
	})

	return appRoute
}
