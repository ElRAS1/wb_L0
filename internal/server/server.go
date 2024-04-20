package server

import (
	"io"
	"net/http"

	"github.com/ElRAS1/wb_L0/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APPServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// Create config
func New(config *Config) *APPServer {
	return &APPServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start server
func (s *APPServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting server...")
	return http.ListenAndServe(s.config.Addr, s.router)
}

func (s *APPServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APPServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APPServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return nil
	}

	s.store = st

	return nil
}

func (s *APPServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
