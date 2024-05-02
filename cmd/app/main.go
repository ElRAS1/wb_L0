package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/ElRAS1/wb_L0/internal/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/app.toml", "path to config file")

}

func main() {
	flag.Parse()

	config := server.NewConfig()

	// Чтение конфигурации из TOML файла
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
