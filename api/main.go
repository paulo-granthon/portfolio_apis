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
	"path/filepath"

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

// generateMarkdown builds a user's portfolio and renders it to a markdown file.
// It accepts two optional flags: -user (defaults to 1) and -out (defaults to
// ../docs/portfolio_<user>.md, i.e. the repository's docs/ folder). It reuses the
// same PortfolioService and RenderMarkdown as the HTTP endpoint, so the file and
// the endpoint produce identical output.
func generateMarkdown(args []string) error {
	fs := flag.NewFlagSet("markdown", flag.ContinueOnError)
	userId := fs.Uint64("user", 1, "id of the user whose portfolio to render")
	out := fs.String("out", "", "output path (defaults to ../docs/portfolio_<user>.md)")
	if err := fs.Parse(args); err != nil {
		return tracerr.Errorf("failed to parse markdown flags: %w", tracerr.Wrap(err))
	}

	outPath := *out
	if outPath == "" {
		outPath = fmt.Sprintf("../docs/portfolio_%d.md", *userId)
	}

	db, err := storage.NewPostgreStorage()
	if err != nil {
		return tracerr.Errorf("failed to create database: %w", tracerr.Wrap(err))
	}

	svc, err := service.NewService(db)
	if err != nil {
		return tracerr.Errorf("failed to create service: %w", tracerr.Wrap(err))
	}

	portfolio, err := svc.PortfolioService.Build(*userId)
	if err != nil {
		return tracerr.Errorf("failed to build portfolio: %w", tracerr.Wrap(err))
	}

	markdown := service.RenderMarkdown(*portfolio)

	if dir := filepath.Dir(outPath); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return tracerr.Errorf("failed to create output directory %q: %w", dir, tracerr.Wrap(err))
		}
	}

	if err := os.WriteFile(outPath, []byte(markdown), 0o644); err != nil {
		return tracerr.Errorf("failed to write markdown to %q: %w", outPath, tracerr.Wrap(err))
	}

	fmt.Printf("Wrote portfolio markdown to %s\n", outPath)
	return nil
}
