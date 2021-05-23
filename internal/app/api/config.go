package api

import "github.com/skullkon/go_lab/helpers"

type Config struct {
	Port        string
	LoggerLevel string
}

var (
	port        = helpers.GetEnvDefault("PORT", ":8080")
	loggerLevel = helpers.GetEnvDefault("loggerLevel", "debug")
)

func NewConfig() *Config {

	return &Config{
		Port:        port,
		LoggerLevel: loggerLevel,
	}
}
