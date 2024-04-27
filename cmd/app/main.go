package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	server "github.com/ElRAS1/wb_L0/internal/server"
	"github.com/joho/godotenv"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/app.toml", "path to config file")
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	// Чтение конфигурации из TOML файла
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}

	err = initConfig(config)

	if err != nil {
		log.Fatalln(err)
	}

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}

func initConfig(config *server.Config) error {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" {
		return fmt.Errorf("one or more required environment variables are not set")
	}

	config.Store.DBHost = dbHost
	config.Store.DBName = dbName
	config.Store.DBPassword = dbPassword
	config.Store.DBUser = dbUser

	return nil
}
