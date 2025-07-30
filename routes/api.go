package routes

import (
	_ "github.com/agastiya/tiyago/docs"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (routes *Routes) InitRoutes() *chi.Mux {

	r := chi.NewRouter()
	rm := routes.Middleware
	rc := routes.Controller

	r.Group(func(r chi.Router) {
		r.Use(chiMiddleware.RequestID)
		r.Use(chiMiddleware.RealIP)
		r.Use(chiMiddleware.RedirectSlashes)
		r.Use(chiMiddleware.Recoverer)

		r.Route("/tiyago", func(r chi.Router) {
			r.Get("/ping", rc.BaseController.Ping)
			r.Route("/auth", func(r chi.Router) {
				r.Post("/loginbyemail", rc.AuthController.LoginByEmail)
				r.Post("/refreshtoken", rc.AuthController.RefreshToken)
			})
			r.Group(func(r chi.Router) {
				r.Use(rm.UserAuth())
				r.Route("/users", func(r chi.Router) {
					r.With(rm.Guard("user.browse")).Get("/", rc.UserController.UserBrowse)
					r.With(rm.Guard("user.create")).Post("/", rc.UserController.UserCreate)
					r.Route("/{id}", func(r chi.Router) {
						r.With(rm.Guard("user.detail")).Get("/", rc.UserController.UserDetail)
						r.With(rm.Guard("user.update")).Put("/", rc.UserController.UserUpdate)
						r.With(rm.Guard("user.updatePassword")).Put("/password", rc.UserController.UserUpdatePassword)
						r.With(rm.Guard("user.delete")).Delete("/", rc.UserController.UserDelete)
					})
				})
			})
			switch routes.Env {
			case "local":
				r.Mount("/swagger", httpSwagger.WrapHandler)
			case "development":
				r.Group(func(r chi.Router) {
					r.With(rm.BasicAuthSwagger()).Mount("/swagger", httpSwagger.WrapHandler)
				})
			}
		})
	})

	return r
}
