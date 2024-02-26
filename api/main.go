package main

import (
	"log"
	"server"
)

func main() {
	server, error := server.NewServer(
		3333,
		CreateEndpoints(),
	)

	if error != nil {
		log.Fatal("Error creating server:", error)
		return
	}

	server.Start()
}
