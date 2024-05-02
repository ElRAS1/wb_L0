package store

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Store struct {
	config     *Config
	db         *sqlx.DB
	Cache      map[string]interface{}
	CacheMutex sync.Mutex
}

func New(config *Config) *Store {
	return &Store{
		config: config,
		Cache:  make(map[string]interface{}),
	}
}
func (s *Store) Open() error {
	err := initConfig(s.config)
	if err != nil {
		log.Fatalln(err)
	}
	DatabaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", s.config.DBName, s.config.DBPassword, s.config.DBHost, s.config.DBPort, s.config.DBName)
	db, err := sqlx.Connect("postgres", DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func initConfig(config *Config) error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" {
		return fmt.Errorf("one or more required environment variables are not set")
	}

	config.DBHost = dbHost
	config.DBUser = dbUser
	config.DBPassword = dbPassword
	config.DBName = dbName
	config.DBPort = dbPort

	return nil
}
