package server

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	server *http.Server
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

	s.server = &http.Server{
		Addr:    s.config.Addr,
		Handler: s.router,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		s.logger.Info("Received shutdown signal")
		s.server.Close()
	}()

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
	ns, err := stan.Connect("test-cluster", "order", stan.NatsURL("nats:4222"))

	if err != nil {
		return err
	}

	defer ns.Close()
	s.store.NatsSubscribe(ns)

	s.store.NatsPublish(ns)

	return nil
}

// func (s *APPServer) getOrder(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	if res, ok := s.store.Cache[id]; ok {
// 		js, err := json.MarshalIndent(res, "", " ")
// 		if err != nil {
// 			http.Error(w, "internal server error", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		s.logger.Info("Returning the response to the GET request")
// 		_, err = w.Write([]byte(js))
// 		if err != nil {
// 			http.Error(w, "data could not be published", http.StatusUnprocessableEntity)
// 		}
// 	} else {
// 		s.logger.Error("no data available")
// 		http.Error(w, "data not found", http.StatusNotFound)
// 	}

// }

func (s *APPServer) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	res, ok := s.store.Cache[id]
	if !ok {
		s.logger.Error("no data available")
		http.Error(w, "data not found", http.StatusNotFound)
		return
	}

	js, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s.logger.Info("Returning the response to the GET request")
	_, err = w.Write([]byte(js))
	if err != nil {
		http.Error(w, "data could not be published", http.StatusUnprocessableEntity)
	}

}
