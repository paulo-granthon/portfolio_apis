package main

import (
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
	router.HandleFunc("/", s.getRoot).Methods("GET")

	return http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
