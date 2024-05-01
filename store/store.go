package store

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
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

	DatabaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", s.config.DBName, s.config.DBPassword, s.config.DBHost, s.config.DBName)
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
