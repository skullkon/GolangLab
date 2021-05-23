package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (a *API) configLogger() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configRouter() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
}

func (a *API) configureStorage() error {
	storage := storage.New(a.config.Storage)
}
