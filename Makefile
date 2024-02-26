FRONTEND_DIR = web
BACKEND_DIR = api

### Whole project commands

# default execution when calling `make`
all: all-dev

# Run all the setup and start both api and web
all-dev: all-setup mono-dev

# Setup the whole project
all-setup:
	@make mono-setup
	@make backend-build
	@make frontend-setup

# Clean the whole project's temporary files
all-clean:
	@make backend-clean
	@make frontend-clean
	@make mono-clean

### Mono Repository only commands

# Run both api and web in development mode
mono-dev:
	@yarn dev

# Run both api and web in production mode
mono-prod:
	@yarn prod

# Setup concurrently wrapper
mono-setup:
	@yarn setup

# Clean concurrently wrapper's temporary files
mono-clean:
	@rm -rf node_modules

### Frontend commands

# Run frontend in development mode
frontend-dev:
	@cd $(FRONTEND_DIR) && yarn dev

# Run frontend in production mode
frontend-prod:
	@make frontend-setup
	@make frontend-build
	@cd $(FRONTEND_DIR) && yarn prod

# Setup frontend dependencies
frontend-setup:
	@cd $(FRONTEND_DIR) && yarn

# Build frontend for production
frontend-build:
	@cd $(FRONTEND_DIR) && yarn build

# Clean frontend's temporary files and production build
frontend-clean:
	@cd $(FRONTEND_DIR) && rm -rf node_modules
	@cd $(FRONTEND_DIR) && rm -rf dist

### Backend commands

# Run backend tests
backend-test:
	@go test -C ./api/ -v

# Run backend in development mode
backend-dev:
	@cd $(BACKEND_DIR) && ./stop.bash && air

# Run backend in production mode
backend-prod:
	@make backend-build
	@./bin/api/main

# Build backend for production
backend-build:
	@go build -C ./api/ -o ../bin/api/main -tags prod

# Clean backend production build
backend-clean:
	@rm -rf bin

### Shortcuts / Aliases

a: all

ad: all-dev
as: all-setup
ac: all-clean

md: mono-dev
mp: mono-prod
ms: mono-setup
mc: mono-clean

bt: backend-test
bd: backend-dev
bp: backend-prod
bs: backend-build
bc: backend-clean

fd: frontend-dev
fp: frontend-prod
fb: frontend-build
fs: frontend-setup
fc: frontend-clean

### Phony targets

.PHONY: \
	all \
	all-dev \
	all-setup \
	all-clean \
	\
	mono-dev \
	mono-prod \
	mono-setup \
	mono-clean \
	\
	backend-test \
	backend-dev \
	backend-prod \
	backend-build \
	backend-clean \
	\
	frontend-dev \
	frontend-prod \
	frontend-setup \
	frontend-build \
	frontend-clean \
	\
	a ad as \
	md mp ms mc \
	fd fp fb fs fc \
	bt bd bp bs bc \
	ac
