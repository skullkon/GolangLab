package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *API) Start() error {
	if err := a.configLogger(); err != nil {
		return err
	}
	a.logger.Info("Starting server at port:", a.config.Port)
	a.configRouter()
	return http.ListenAndServe(a.config.Port, a.router)
}
