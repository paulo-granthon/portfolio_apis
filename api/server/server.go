package server

import (
	"fmt"
	"log"
	"net/http"
	"service"
	"storage"

	"github.com/gorilla/mux"
	"github.com/ztrue/tracerr"

	_ "docs"

	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

type Server struct {
	port      string
	endpoints []Endpoint
	Storage   storage.Storage
	Service   service.Service
}

func NewServer(
	port int,
	endpoints []Endpoint,
	storage storage.Storage,
	service service.Service,
) (*Server, error) {
	if port < 1 || port > 65535 {
		return nil, tracerr.Errorf("Invalid listen address: %v", port)
	}

	return &Server{
		port:      fmt.Sprintf(":%v", port),
		endpoints: endpoints,
		Storage:   storage,
		Service:   service,
	}, nil
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	for _, endpoint := range s.endpoints {

		router.HandleFunc(
			endpoint.Path,
			endpoint.Create(*s),
		// allowing all methods to allow custom response when method is not allowed
		).Methods("OPTIONS", "GET", "POST", "PUT", "DELETE")
	}

	router.
		PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	)).Methods("GET")

	log.Println("Starting server on", s.port)

	err := http.ListenAndServe(s.port, router)
	if err != nil {
		return tracerr.Errorf("Error starting server: %w", tracerr.Wrap(err))
	}

	return nil
}
