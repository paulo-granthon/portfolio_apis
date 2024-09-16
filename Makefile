FRONTEND_DIR = web
BACKEND_DIR = api

### Whole project commands

# default execution when calling `make`
all: all-dev

# Run all the setup and start both api and web
all-dev: all-setup mono-dev

# Run all the setup and start both api and web in production mode
all-prod: all-setup mono-prod

# Setup the whole project
all-setup:
	@make mono-setup
	@make back-build
	@make front-setup

# Clean the whole project's temporary files
all-clean:
	@make back-clean
	@make front-clean
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
front-dev:
	@cd $(FRONTEND_DIR) && yarn dev

# Run frontend in production mode
front-prod:
	@make front-build
	@cd $(FRONTEND_DIR) && yarn prod

# Setup frontend dependencies
front-setup:
	@cd $(FRONTEND_DIR) && yarn

# Build frontend for production
front-build:
	@cd $(FRONTEND_DIR) && yarn build

# Clean frontend's temporary files and production build
front-clean:
	@cd $(FRONTEND_DIR) && rm -rf node_modules
	@cd $(FRONTEND_DIR) && rm -rf dist

### Backend commands

# Run backend tests
back-test:
	@go test -C ./api/ -v

# Run backend in development mode
back-dev:
	@cd $(BACKEND_DIR) && ./stop.bash && air

# Run backend in production mode
back-prod:
	@make back-build
	@./bin/api/main

# Build backend for production
back-build:
	@go build -C ./api/ -o ../bin/api/main -tags prod

# Clean backend production build
back-clean:
	@rm -rf bin

### Database commands

# Start database container
database-up:
	@docker-compose up -d

# Stop database container
database-down:
	@docker-compose down

# Migrate database
database-migrate:
	@make database-up
	@cd $(BACKEND_DIR) && go run main.go seed

# Recreate and start the database container
database-recreate:
	@make database-down
	@docker-compose up --build -d

### Shortcuts / Aliases

a: all

ad: all-dev
as: all-setup
ac: all-clean

md: mono-dev
mp: mono-prod
ms: mono-setup
mc: mono-clean

bt: back-test
bd: back-dev
bp: back-prod
bs: back-build
bc: back-clean

fd: front-dev
fp: front-prod
fb: front-build
fs: front-setup
fc: front-clean

du: database-up
dm: database-migrate
dd: database-down
dr: database-recreate

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
	back-test \
	back-dev \
	back-prod \
	back-build \
	back-clean \
	\
	front-dev \
	front-prod \
	front-setup \
	front-build \
	front-clean \
	\
	a ad as \
	md mp ms mc \
	fd fp fb fs fc \
	bt bd bp bs bc \
	du dm dd dr \
	ac
