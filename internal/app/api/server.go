package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/skullkon/go_lab/storage"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
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
	a.logger.Info("Router has configured successefully")
	if err := a.configStorage(); err != nil {
		return err
	}
	a.logger.Info("Storage has configured successefully")
	return http.ListenAndServe(a.config.Port, a.router)
}
