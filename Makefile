FRONTEND_DIR = web
BACKEND_DIR = api

.PHONY: frontend-dev frontend-build frontend-setup backend-dev

all: setup mono

setup:
	yarn setup

mono:
	yarn dev

fd: frontend-dev
frontend-dev:
	cd $(FRONTEND_DIR) && yarn dev

fb: frontend-build
frontend-build:
	cd $(FRONTEND_DIR) && yarn build

fs: frontend-setup
frontend-setup:
	cd $(FRONTEND_DIR) && yarn

bd: backend-dev
backend-dev:
	cd $(BACKEND_DIR) && go run main.go
