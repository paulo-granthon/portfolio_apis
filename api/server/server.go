package server

import (
	"fmt"
	"log"
	"net/http"
	"storage"

	"github.com/gorilla/mux"
	"github.com/ztrue/tracerr"
)

type Server struct {
	port      string
	endpoints []Endpoint
	Storage   storage.Storage
}

func NewServer(
	port int,
	endpoints []Endpoint,
	storage storage.Storage,
) (*Server, error) {
	if port < 1 || port > 65535 {
		return nil, tracerr.Errorf("Invalid listen address: %v", port)
	}

	return &Server{
		port:      fmt.Sprintf(":%v", port),
		endpoints: endpoints,
		Storage:   storage,
	}, nil
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	for _, endpoint := range s.endpoints {
		methods := []string{}
		for _, method := range endpoint.Methods {
			methods = append(methods, method.Method)
		}

		router.HandleFunc(
			endpoint.Path,
			endpoint.Create(*s),
		// allowing all methods to allow custom response when method is not allowed
		).Methods("OPTIONS", "GET", "POST", "PUT", "DELETE")
	}

	log.Println("Starting server on", s.port)

	err := http.ListenAndServe(s.port, router)
	if err != nil {
		return tracerr.Errorf("Error starting server: %w", tracerr.Wrap(err))
	}

	return nil
}
