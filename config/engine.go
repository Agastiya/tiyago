package config

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/agastiya/tiyago/contracts"
	"github.com/agastiya/tiyago/database/migrations"
	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/module"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/agastiya/tiyago/pkg/jwt"
	"github.com/agastiya/tiyago/routes"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type (
	Env dto.Environment

	Config struct {
		Environment Env
		Engine      contracts.Engine
	}
)

func GetEnvironment(env string) *Config {

	var environment Env

	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), constant.Environment+env+".yml")

	_, err := os.Stat(envPath)
	if err != nil {
		log.Fatal().Msgf("Environment file not found: %s", envPath)
	}

	content, err := os.ReadFile(envPath)
	if err != nil {
		log.Fatal().Msgf("Failed to read file. %+v", err.Error())
	}

	err = yaml.Unmarshal(content, &environment)
	if err != nil {
		log.Fatal().Msgf("Failed to Unmarshal. %+v", err.Error())
	}

	log.Info().Msgf("[%s] Environment Configuration Loaded Successfully!", environment.App.Environment)
	return &Config{Environment: environment, Engine: environment}
}

func (env Env) InitDatabase() {
	dbConfig := env.Databases[0]
	connectionName := DBConfigName(env.Databases[0].Connection)
	dbConfigConnection[connectionName] = CreatePostgreSQLConnection(dbConfig)
}

func (env Env) Migrate() {
	db := DATABASE_MAIN.Get()
	migrations.Run(db)
	os.Exit(1)
}

func (env Env) InitPackage() {
	jwt.JwtVar = &jwt.JwtService{
		ConfigJwt: env.Jwt,
	}
}

func (env Env) InitRoute() *chi.Mux {
	db := DATABASE_MAIN.Get()
	routes := &routes.Routes{
		Env:        env.App.Environment,
		Controller: module.InitModules(db).Controller,
		// Middleware: &Middleware.Middleware{
		// 	Jwt:            Jwt.JwtVar,
		// 	SwaggerSetting: environment.Environment.Swagger,
		// },
	}
	return routes.InitRoutes()
}

func (env Env) Serve() {

	var healthy int32
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	address := env.App.Host + ":" + env.App.Port
	httpServer := &http.Server{
		Addr:    address,
		Handler: env.InitRoute(),
	}

	go func() {
		<-quit
		log.Info().Msg("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal().Msgf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Info().Msgf("Http Service running on %s %s", address, " .....")
	atomic.StoreInt32(&healthy, 1)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Could not listen on %s: %v\n", address, err)
	}

	<-done
	log.Info().Msg("Server stopped")
}
