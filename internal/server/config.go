package server

import (
	"github.com/ElRAS1/wb_L0/store"
)

type Config struct {
	Addr     string `yaml:"addr"`
	LogLevel string `yaml:"loglevel"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
