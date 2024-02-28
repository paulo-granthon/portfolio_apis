package server

import (
	"fmt"
	"log"
	"net/http"
	"storage"

	"github.com/gorilla/mux"
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
		return nil, fmt.Errorf("Invalid listen address: %v", port)
	}

	return &Server{
		port:      fmt.Sprintf(":%v", port),
		endpoints: endpoints,
		Storage:   storage,
	}, nil
}

func (s *Server) Start() {
	router := mux.NewRouter()

	for _, endpoint := range s.endpoints {
		methods := []string{}
		for _, method := range endpoint.Methods {
			methods = append(methods, method.Method)
		}
		router.HandleFunc(endpoint.Path, endpoint.Create(*s)).Methods(methods...)
	}

	log.Println("Starting server on", s.port)

	error := http.ListenAndServe(s.port, router)
	if error != nil {
		log.Fatal("Error starting server:", error)
	}
}
