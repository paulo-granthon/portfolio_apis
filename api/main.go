package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint().Create()).Methods("GET")

	log.Println("Starting server on", s.listenAddress)

	return http.ListenAndServe(s.listenAddress, router)
}
