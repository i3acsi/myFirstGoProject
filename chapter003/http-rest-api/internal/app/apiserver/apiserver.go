package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
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

func (s *APIServer) handleHello() http.HandlerFunc {
	// define some args for this handler
	// this code 'll done once, on init
	// example: type req struct{ name string }

	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello!")
	}
}
