package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/coinprofile/services/gotemplate/src/config"
	"gitlab.com/coinprofile/services/gotemplate/src/handlers"
	"gitlab.com/coinprofile/services/gotemplate/src/middlewares"
	"gitlab.com/coinprofile/services/gotemplate/src/services"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	gorillaHandlers "github.com/gorilla/handlers"
)

// Run starts the server
func Run() {
	cfg := config.Init()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	r := mux.NewRouter().StrictSlash(true)
	h := routing(cfg, r)
	PORT, _ := strconv.Atoi(cfg.Port)
	log.Infof("Application [ %s ] started on port [ %s ]", cfg.AppName, cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), h))
}

func routing(cfg *config.Config, r *mux.Router) http.Handler {

	r.HandleFunc("/ping", handlers.Handler(cfg, &services.PingReq{})).Methods(http.MethodGet)

	r.Use(middlewares.ResponseHeaderMiddleware)
	r.Use(middlewares.RecoverMiddleware)

	loggedRouter := gorillaHandlers.LoggingHandler(os.Stdout, r)

	allowedOrigins := []string{"*"}
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	corsRouter := corsOpts.Handler(loggedRouter)
	return corsRouter
}
