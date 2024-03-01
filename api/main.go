package main

import (
	"endpoints"
	"log"
	"server"
	"storage"
)

func main() {
	db, error := storage.NewPostgreStorage()
	if error != nil {
		log.Fatal("Error creating database: ", error)
		return
	}

	server, error := server.NewServer(
		3333,
		endpoints.CreateEndpoints(),
		db,
	)

	if error != nil {
		log.Fatal("Error creating server: ", error)
		return
	}

	server.Start()
}
