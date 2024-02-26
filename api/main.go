package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(port int) *APIServer {
	return &APIServer{
		listenAddress: fmt.Sprintf(":%v", port),
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint().Create()).Methods("GET")

	log.Println("Starting server on", s.listenAddress)

	return http.ListenAndServe(s.listenAddress, router)
func main() {
	server := NewAPIServer(3333)
	server.Start()
}
