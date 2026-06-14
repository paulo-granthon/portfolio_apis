# web

React + TypeScript frontend. It fetches the composed portfolio document from the
`api` backend (`GET /portfolio/{id}`) and renders it interactively, and links to
the markdown export for download. Built with Vite, StyleX and React Router.

The app has a single route: `/portfolio/:userId` (e.g. `http://localhost:5173/portfolio/1`).

## Prerequisites

- **Node 18+** and Yarn (the repo uses Yarn 4 workspaces)
- Install deps once at the **repository root**: `yarn install`
- The `api` backend running on `:3333` (see `../api/README.md`) for live data

> All `nx` commands below run from the **repository root**. Prefix with `yarn`
> (`yarn nx ...`) if `nx` isn't installed globally.

## Configuration

Environment variables live in `web/.env` (Vite exposes `VITE_`-prefixed vars):

| Variable            | Example                  | Notes                                            |
| ------------------- | ------------------------ | ------------------------------------------------ |
| `VITE_API_URL`      | `http://localhost:3333`  | base URL of the `api` backend                    |
| `VITE_GITHUB_TOKEN` | `your_github_token`      | optional, used to fetch GitHub profile pictures  |

## Run

```sh
nx serve web          # start the Vite dev server (default http://localhost:5173)
nx dev web            # alias — runs `vite` directly
```

## Build & preview

```sh
nx build web          # type-check (tsc) + production build into web/dist
nx preview web        # serve the production build locally
nx run web:serve-static  # serve the already-built web/dist
```

## Quality

```sh
nx test web           # run the Vitest + React Testing Library suite (jsdom)
nx typecheck web      # tsc --noEmit
nx lint web           # eslint
```

## Project layout

```
web/src/
├── pages/        portfolio page (fetches the document, renders + download link)
├── components/   presentational React components (User, ProjectList, Project, ...)
├── services/     API clients (portfolio, user, project, contribution, github)
├── schemas/      TypeScript types mirroring the api models
└── styles/       StyleX styles
```

The data flow is one fetch: `services/portfolio.ts` calls `/portfolio/{id}` and
maps it into typed schemas; the components render from props (no per-component
fetching). The "Download Markdown" link points at `/portfolio/{id}/markdown`.

## Running api + web together

From the repo root:

```sh
nx run-many -t serve -p api web
```
