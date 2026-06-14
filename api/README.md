# api

Go backend for the portfolio project. It maps the database, serves the JSON the
`web` app renders, and is the single source of truth for the portfolio document
(both the structured `/portfolio/{id}` endpoint and the markdown export derive
from the same model).

Single Go module: `github.com/paulo-granthon/portfolio_apis`. The server listens
on port **3333**; Swagger UI is served at `/swagger`.

## Prerequisites

- **Go 1.25+**
- **Docker** (for the local PostgreSQL database)
- Two Go CLI tools, used by the `swagger` and `dev` targets:
  ```sh
  go install github.com/swaggo/swag/cmd/swag@latest
  go install github.com/air-verse/air@latest
  ```
- Node deps installed once at the repo root (`yarn install`) so the `nx` CLI is
  available.

> All `nx` commands below are run from the **repository root**. Prefix with
> `yarn` (`yarn nx ...`) if you don't have `nx` installed globally.

## Configuration

Connection settings live in `api/.env`:

| Variable  | Default     | Notes                         |
| --------- | ----------- | ----------------------------- |
| `DB_HOST` | `0.0.0.0`   |                               |
| `DB_PORT` | `3332`      | host port mapped to Postgres  |
| `DB_USER` | `postgres`  |                               |
| `DB_PASS` | `secret`    |                               |
| `DB_NAME` | (empty)     | uses the default database     |

## Database

```sh
nx run api:db-up      # start the PostgreSQL container (built from ./Dockerfile, includes the pguint extension)
nx run api:db-down    # stop and remove it
nx run api:seed       # drop, recreate and populate every table (users, teams, projects, skills, contributions)
```

`seed` is destructive: it recreates the schema from scratch before inserting the
seed data in `seeds/seeds.go`.

## Run

```sh
nx serve api          # go run . (regenerates Swagger first); good for a one-off run
nx run api:dev        # live-reload via Air (swag init + rebuild on change); preferred for development
```

Both boot the HTTP server on `:3333`. The database must be up and seeded first.

## Build

```sh
nx build api          # regenerates Swagger, then builds the binary to dist/api
```

## Quality

```sh
nx test api           # go test ./...
nx lint api           # gofmt check
nx tidy api           # go mod tidy
```

## Swagger / API docs

```sh
nx run api:swagger    # runs `swag init`, generating api/docs (gitignored)
```

`build`, `serve` and `seed` depend on this target, so docs are regenerated
automatically before they run. Once the server is up, browse the UI at
`http://localhost:3333/swagger`.

## Portfolio markdown export

The portfolio document can be produced two ways, from the same renderer (their
output is byte-identical):

- **HTTP**: `GET /portfolio/{id}/markdown` returns it as a download.
- **CLI / target**:
  ```sh
  nx run api:markdown                       # user 1 -> docs/portfolio_1.md (defaults)
  nx run api:markdown --user=2              # user 2 -> docs/portfolio_2.md
  nx run api:markdown --out=/tmp/p.md       # user 1 -> /tmp/p.md
  nx run api:markdown --user=2 --out=p.md   # both overridden
  ```
  Both flags are optional: `--user` defaults to **1**, and `--out` defaults to
  the repository's **`docs/portfolio_<user>.md`**. The output rendered here is
  byte-identical to the HTTP endpoint, and includes the user's GitHub avatar.

## Running the binary directly

Inside `api/`, the compiled program also accepts subcommands (this is what the
`seed` and `markdown` targets wrap):

```sh
go run . seed                              # seed the database
go run . markdown [-user=1] [-out=path]    # render a portfolio (defaults: user 1, ../docs/portfolio_<user>.md)
go run .                                   # start the server
```

## Running api + web together

From the repo root:

```sh
nx run-many -t serve -p api web
```
