package api

import (
	"github.com/skullkon/go_lab/helpers"
	"github.com/skullkon/go_lab/storage"
)

type Config struct {
	Port        string
	LoggerLevel string
	Storage     *storage.Config
}

var (
	port        = helpers.GetEnvDefault("PORT", ":8080")
	loggerLevel = helpers.GetEnvDefault("loggerLevel", "debug")
)

func NewConfig() *Config {

	return &Config{
		Port:        port,
		LoggerLevel: loggerLevel,
		Storage:     storage.NewConfig(),
	}
}
