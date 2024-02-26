package main

import (
	"endpoints"
	"log"
	"server"
)

func main() {
	server, error := server.NewServer(
		3333,
		endpoints.CreateEndpoints(),
	)

	if error != nil {
		log.Fatal("Error creating server:", error)
		return
	}

	server.Start()
}
