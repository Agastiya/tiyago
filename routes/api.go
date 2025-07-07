package routes

import (
	"github.com/agastiya/tiyago/controller"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Routes struct {
	Env string
	// Middleware Middleware.MiddlewareInterface
	Controller controller.ControllerInterface
}

func (app *Routes) CollectRoutes() *chi.Mux {

	appRoute := chi.NewRouter()
	appRoute.Use(chiMiddleware.RequestID)
	appRoute.Use(chiMiddleware.RealIP)
	appRoute.Use(chiMiddleware.RedirectSlashes)
	appRoute.Use(chiMiddleware.Recoverer)

	appRoute.Route("/tiyago", func(appRoute chi.Router) {

		// appRoute.Group(func(appRoute chi.Router) {
		// 	appRoute.Route("/auth", func(appRoute chi.Router) {
		// 		appRoute.Post("/login", app.Controller.Login)
		// 	})
		// })

		// appRoute.Group(func(appRoute chi.Router) {
		// 	appRoute.Use(app.Middleware.UserAuth())
		// 	appRoute.Route("/account", func(appRoute chi.Router) {
		// 		appRoute.Post("/", app.Controller.CreateAccount)
		// 		appRoute.Route("/{id}", func(appRoute chi.Router) {
		// 			appRoute.Put("/", app.Controller.UpdateAccountPassword)
		// 			appRoute.Put("/active", app.Controller.UpdateAccountStatus)
		// 		})
		// 	})
		// })

		// use basic auth or mount swagger based on specific environment
		// if app.Env == "development" {
		// 	appRoute.Group(func(appRoute chi.Router) {
		// 		appRoute.Use(app.Middleware.BasicAuthSwagger())
		// 		appRoute.Mount("/swagger", httpSwagger.WrapHandler)
		// 	})
		// } else if app.Env == "local" {
		// 	appRoute.Mount("/swagger", httpSwagger.WrapHandler)
		// }

		appRoute.Get("/ping", app.Controller.Ping)
	})

	return appRoute
}
