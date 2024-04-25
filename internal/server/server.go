package server

import (
	"encoding/json"

	"net/http"

	"github.com/ElRAS1/wb_L0/store"

	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
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

func (s *APPServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	defer s.store.Close()

	s.logger.Info("Connecting  Database...")

	if err := s.configureNats(); err != nil {
		s.logger.Error("Failed connection on nats...")
		return err
	}
	s.logger.Info("Connecting  nats-streaming...")

	s.logger.Info("Starting server... ports: ", s.config.Addr)
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
	s.router.HandleFunc("/order/{id}", s.getOrder)
}

func (s *APPServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APPServer) configureNats() error {
	ns, err := stan.Connect("test-cluster", "test-cluster")

	if err != nil {
		return err
	}
	defer ns.Close()

	s.store.NatsSubscribe(ns)

	s.store.NatsPublish(ns)
	return nil
}

func (s *APPServer) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if res, ok := s.store.Cache[id]; ok {
		js, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		s.logger.Info("Returning the response to the GET request")
		w.Write([]byte(js))
	} else {
		s.logger.Error("no data available")
		http.Error(w, "data not found", http.StatusNotFound)
	}

}
