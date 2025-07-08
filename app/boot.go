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

var (
	environment = flag.String("tag", "", "define tag")
	migrate     = flag.Bool("migrate", false, "run migration only")
	initConfig  *config.Config
)

func init() {
	flag.Parse()

	switch {
	case strings.Contains(*environment, "DEV-"):
		*environment = constant.Development
	case strings.Contains(*environment, "PROD-"):
		*environment = constant.Production
	default:
		*environment = constant.Local
	}

	initConfig = config.GetEnvironment(*environment)
}

func PackageInit(env *config.Environment) {
	jwt.JwtVar = &jwt.JwtService{ConfigJwt: env.Jwt}
}

func AppInit() {

	initConfig.Engine.BuildConnection()
	initConfig.Engine.RunMigration(migrate)
	PackageInit(&initConfig.Environment)
	initConfig.Routes = &routes.Routes{
		Env:        initConfig.Environment.App.Environment,
		Controller: &controller.Controller{},
		// Middleware: &Middleware.Middleware{
		// 	Jwt:            Jwt.JwtVar,
		// 	SwaggerSetting: environment.Environment.Swagger,
		// },
	}
	initConfig.Engine.ServeHTTP(initConfig.Routes.CollectRoutes())
}
