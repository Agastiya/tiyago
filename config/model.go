package config

import "github.com/go-chi/chi/v5"

type (
	Config struct {
		Database    DatabaseInterface
		Routes      RouteInterface
		Engine      EngineInterface
		Environment Environment
	}

	DatabaseInterface interface {
		BuildConnection()
	}

	EngineInterface interface {
		ServeHTTP(route *chi.Mux)
	}

	RouteInterface interface {
		CollectRoutes() *chi.Mux
	}

	Environment struct {
		App       AppSetting        `yaml:"app"`
		Databases []DatabaseSetting `yaml:"databases"`
		Jwt       JwtSetting        `yaml:"jwt"`
		Swagger   SwaggerSetting    `yaml:"swagger"`
	}

	AppSetting struct {
		Name        string `yaml:"name"`
		Service     string `yaml:"service"`
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Environment string `yaml:"environment"`
	}

	DatabaseSetting struct {
		Connection string `yaml:"connection"`
		Host       string `yaml:"host"`
		Port       int64  `yaml:"port"`
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		Usage      string `yaml:"usage"`
	}

	JwtSetting struct {
		SecretKey string `yaml:"secretKey"`
	}

	SwaggerSetting struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)
