package main

import (
	"endpoints"
	"flag"
	"fmt"
	"seeds"
	"server"
	"storage"

	"github.com/ztrue/tracerr"
)

func main() {
	if handleArgs() {
		return
	}

	db, err := storage.NewPostgreStorage()
	if err != nil {
		tracerr.PrintSourceColor(
			tracerr.Errorf("Error creating database: %w", tracerr.Wrap(err)),
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
			tracerr.Errorf("Error creating server: %w", tracerr.Wrap(err)),
		)
		return
	}

	if err := server.Start(); err != nil {
		tracerr.PrintSourceColor(
			tracerr.Errorf("Error starting server: %w", tracerr.Wrap(err)),
		)
		return
	}
}

func handleArgs() bool {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		return false
	}

	switch args[0] {
	case "seed":
		err := seeds.Run()
		if err != nil {
			tracerr.PrintSourceColor(
				tracerr.Errorf("Error seeding database: %w", tracerr.Wrap(err)),
			)
		}

		fmt.Println("Seeding complete")
	default:
		fmt.Printf("Unrecognized flag %v\n", args[0])
	}

	return true
}
