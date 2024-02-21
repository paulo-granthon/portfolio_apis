FRONTEND_DIR = web
BACKEND_DIR = api

.PHONY: frontend-dev frontend-build frontend-setup backend-dev

fd: frontend-dev
frontend-dev:
	cd $(FRONTEND_DIR) && yarn start

fb: frontend-build
frontend-build:
	cd $(FRONTEND_DIR) && yarn build

fs: frontend-setup
frontend-setup:
	cd $(FRONTEND_DIR) && yarn

bd: backend-dev
backend-dev:
	cd $(BACKEND_DIR) && go run main.go


