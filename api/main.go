package main

import (
	"endpoints"
	"server"
	"storage"

	"github.com/ztrue/tracerr"
)

func main() {
	db, err := storage.NewPostgreStorage()
	if err != nil {
		tracerr.PrintSourceColor(
			tracerr.Errorf("Error creating database: :%w", tracerr.Wrap(err)),
		)
		return
	}

	server, err := server.NewServer(
		3333,
		endpoints.CreateEndpoints(),
		db,
	)
	if err != nil {
		tracerr.PrintSourceColor(
			tracerr.Errorf("Error creating server: :%w", tracerr.Wrap(err)),
		)
		return
	}

	if err := server.Start(); err != nil {
		tracerr.PrintSourceColor(
			tracerr.Errorf("Error starting server: :%w", tracerr.Wrap(err)),
		)
		return
	}
}
