package routes

import (
	_ "github.com/agastiya/tiyago/docs"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (routes *Routes) InitRoutes() *chi.Mux {

	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.RedirectSlashes)
	r.Use(chiMiddleware.Recoverer)

	r.Route("/tiyago", func(r chi.Router) {
		r.Get("/ping", routes.Controller.BaseController.Ping)
		r.Route("/auth", func(r chi.Router) {
			r.Post("/loginbyemail", routes.Controller.AuthController.LoginByEmail)
			r.Post("/refreshtoken", routes.Controller.AuthController.RefreshToken)
		})
		r.Group(func(r chi.Router) {
			r.Use(routes.Middleware.UserAuth())
			r.Route("/user", func(r chi.Router) {
				r.Get("/", routes.Controller.UserController.UserBrowse)
				r.Post("/", routes.Controller.UserController.UserCreate)
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", routes.Controller.UserController.UserDetail)
					r.Put("/", routes.Controller.UserController.UserUpdate)
					r.Put("/password", routes.Controller.UserController.UserUpdatePassword)
					r.Delete("/", routes.Controller.UserController.UserDelete)
				})
			})
		})
		switch routes.Env {
		case "local":
			r.Mount("/swagger", httpSwagger.WrapHandler)
		case "development":
			r.Group(func(r chi.Router) {
				r.With(routes.Middleware.BasicAuthSwagger()).Mount("/swagger", httpSwagger.WrapHandler)
			})
		}
	})

	return r
}
