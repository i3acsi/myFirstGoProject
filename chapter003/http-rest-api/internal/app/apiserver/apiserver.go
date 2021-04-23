package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"src/myFirstGoProject/chapter003/http-rest-api/internal/app/store"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server")
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	s.logger.SetLevel(level)
	return nil
}

// configureRouter ...
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// configureStore ...
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// define some args for this handler
	// this code 'll done once, on init
	// example: type req struct{ name string }

	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello!")
	}
}
