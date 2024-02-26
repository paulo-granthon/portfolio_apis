package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(port int) (*APIServer, error) {
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("Invalid listen address: %v", port)
	}

	return &APIServer{
		listenAddress: fmt.Sprintf(":%v", port),
	}, nil
}

func (s *APIServer) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint().Create()).Methods("GET")

	log.Println("Starting server on", s.listenAddress)

	error := http.ListenAndServe(s.listenAddress, router)
	if error != nil {
		log.Fatal("Error starting server:", error)
	}
}

func main() {
	server, error := NewAPIServer(3333)

	if error != nil {
		log.Fatal("Error creating server:", error)
		return
	}

	server.Start()
}
