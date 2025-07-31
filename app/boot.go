package app

import (
	"flag"
	"os"
	"strings"

	"github.com/agastiya/tiyago/config"
	"github.com/agastiya/tiyago/pkg/constant"
)

var (
	environment = flag.String("tag", "", "define tag")
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

func AppInit() {
	initConfig.Engine.InitDatabase()

	if *environment == constant.Local && len(os.Args) > 1 {
		initConfig.Engine.InitCommand()
	}

	initConfig.Engine.Serve()
}
