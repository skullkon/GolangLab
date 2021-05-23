package storage

import "github.com/skullkon/go_lab/helpers"

type Config struct {
	DatabaseURI string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURI: helpers.GetEnvDefault("database_uri", ""),
	}
}
