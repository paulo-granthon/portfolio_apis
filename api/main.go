package main

import (
	"flag"
	"fmt"
	"github.com/paulo-granthon/portfolio_apis/endpoints"
	"github.com/paulo-granthon/portfolio_apis/seeds"
	"github.com/paulo-granthon/portfolio_apis/server"
	"github.com/paulo-granthon/portfolio_apis/service"
	"github.com/paulo-granthon/portfolio_apis/storage"
	"os"
	"strconv"

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

	service, err := service.NewService(db)
	if err != nil {
		tracerr.PrintSourceColor(
			tracerr.Errorf("Error creating service: %w", tracerr.Wrap(err)),
		)
		return
	}

	server, err := server.NewServer(
		3333,
		endpoints.CreateEndpoints(),
		db,
		*service,
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
	case "markdown":
		if err := generateMarkdown(args[1:]); err != nil {
			tracerr.PrintSourceColor(
				tracerr.Errorf("Error generating markdown: %w", tracerr.Wrap(err)),
			)
		}
	default:
		fmt.Printf("Unrecognized flag %v\n", args[0])
	}

	return true
}

// generateMarkdown builds a user's portfolio and renders it to markdown, writing
// to the optional output path (or stdout). It reuses the same PortfolioService
// and RenderMarkdown as the HTTP endpoint, so the file and the endpoint match.
func generateMarkdown(args []string) error {
	if len(args) < 1 {
		return tracerr.Errorf("usage: markdown <userId> [outPath]")
	}

	userId, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return tracerr.Errorf("invalid userId %q: %w", args[0], tracerr.Wrap(err))
	}

	db, err := storage.NewPostgreStorage()
	if err != nil {
		return tracerr.Errorf("failed to create database: %w", tracerr.Wrap(err))
	}

	svc, err := service.NewService(db)
	if err != nil {
		return tracerr.Errorf("failed to create service: %w", tracerr.Wrap(err))
	}

	portfolio, err := svc.PortfolioService.Build(userId)
	if err != nil {
		return tracerr.Errorf("failed to build portfolio: %w", tracerr.Wrap(err))
	}

	markdown := service.RenderMarkdown(*portfolio)

	if len(args) >= 2 {
		if err := os.WriteFile(args[1], []byte(markdown), 0o644); err != nil {
			return tracerr.Errorf("failed to write markdown to %q: %w", args[1], tracerr.Wrap(err))
		}
		fmt.Printf("Wrote portfolio markdown to %s\n", args[1])
		return nil
	}

	fmt.Print(markdown)
	return nil
}
