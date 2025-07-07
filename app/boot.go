package app

import (
	"flag"
	"strings"

	"github.com/agastiya/tiyago/config"
	"github.com/agastiya/tiyago/controller"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/jwt"
	"github.com/agastiya/tiyago/routes"
)

var environment = flag.String("tag", "", "define tag")

func init() {
	flag.Parse()

	if strings.Contains(*environment, "DEV-") {
		*environment = constant.Development
	} else if strings.Contains(*environment, "PROD-") {
		*environment = constant.Production
	} else {
		*environment = constant.Local
	}
}

func ServiceInit(env *config.Environment) {
	jwt.JwtVar = &jwt.JwtService{ConfigJwt: env.Jwt}
}

func AppInitialization() {
	config := config.GetEnvironment(*environment)
	ServiceInit(&config.Environment)

	config.Database.BuildConnection()
	config.Routes = &routes.Routes{
		Env:        config.Environment.App.Environment,
		Controller: &controller.Controller{},
		// Middleware: &Middleware.Middleware{
		// 	Jwt:            Jwt.JwtVar,
		// 	SwaggerSetting: environment.Environment.Swagger,
		// },
	}
	config.Engine.ServeHTTP(config.Routes.CollectRoutes())
}
