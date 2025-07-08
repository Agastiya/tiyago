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

	"github.com/agastiya/tiyago/pkg/constant"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func GetEnvironment(env string) *Config {

	var environment Environment
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), constant.Environment+env+".yml")

	_, err := os.Stat(envPath)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}

	content, err := os.ReadFile(envPath)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(content, &environment)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}

	log.Info().Msgf("[%s] Environment Configuration Loaded Successfully!", environment.App.Environment)
	return &Config{Environment: environment, Engine: environment, Routes: nil}
}

func (env Environment) BuildConnection() {
	dbConfig := env.Databases[0]
	connectionName := DBConfigName(env.Databases[0].Connection)
	dbConfigConnection[connectionName] = CreatePostgreSQLConnection(dbConfig)
}

func (env Environment) ServeHTTP(route *chi.Mux) {
	var healthy int32
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	address := env.App.Host + ":" + env.App.Port
	httpServer := &http.Server{
		Addr:    address,
		Handler: route,
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
