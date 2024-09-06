package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/oktapascal/go-simpro/app/login"
	"github.com/oktapascal/go-simpro/app/user"
	"github.com/oktapascal/go-simpro/app/welcome"
	"github.com/oktapascal/go-simpro/config"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	config.InitConfig()
	log := config.CreateLoggers(nil)
	validate := config.CreateValidator()
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Timeout(60 * time.Second))
	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.RecoverMiddleware)

	welcomeHandler := welcome.Wire()

	router.Get("/", welcomeHandler.Welcome())
	router.NotFound(welcomeHandler.NotFoundApi())
	router.MethodNotAllowed(welcomeHandler.MethodNotAllowedApi())

	user.Wire(validate, db).InitializeRoute(router)
	login.Wire(validate, db).InitializeRoute(router)

	log.Info(fmt.Sprintf("%s Application Started on http://localhost:%s", viper.GetString("APP_NAME"), viper.GetString("APP_PORT")))
	err = http.ListenAndServe(":"+viper.GetString("APP_PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
