#`Makefile`

.PHONY: install dev build test clean deploy docker k8s security audit

# Build variables
BINARY_NAME=geminizer-enterprise
DOCKER_REGISTRY=registry.geminizer.com
VERSION=$(shell git describe --tags --always --dirty)

# Installation
install:
	@echo "Installing Geminizer Enterprise..."
	cd backend && go mod download
	cd frontend && npm install
	@echo "✅ Installation complete"

# Development
dev:
	@echo "Starting development environment..."
	docker-compose -f infrastructure/docker/docker-compose.dev.yml up -d

dev-backend:
	cd backend && air -c .air.toml

dev-frontend:
	cd frontend && npm run dev

# Building
build:
	@echo "Building Geminizer Enterprise v$(VERSION)..."
	cd backend && go build -o ../bin/$(BINARY_NAME) cmd/server/main.go
	cd frontend && npm run build
	@echo "✅ Build complete"

build-docker:
	@echo "Building Docker images..."
	docker build -t $(DOCKER_REGISTRY)/geminizer-backend:$(VERSION) -f infrastructure/docker/Dockerfile.backend .
	docker build -t $(DOCKER_REGISTRY)/geminizer-frontend:$(VERSION) -f infrastructure/docker/Dockerfile.frontend .

# Testing
test:
	@echo "Running tests..."
	cd backend && go test ./... -v
	cd frontend && npm test
	@echo "✅ Tests complete"

test-security:
	@echo "Running security audit..."
	go run honnef.co/go/tools/cmd/staticcheck ./backend/...
	npm audit --audit-level high
	gosec ./backend/...

# Deployment
deploy: build test-security
	@echo "Deploying Geminizer Enterprise v$(VERSION)..."
	docker push $(DOCKER_REGISTRY)/geminizer-backend:$(VERSION)
	docker push $(DOCKER_REGISTRY)/geminizer-frontend:$(VERSION)
	kubectl set image deployment/geminizer-backend geminizer-backend=$(DOCKER_REGISTRY)/geminizer-backend:$(VERSION)
	kubectl set image deployment/geminizer-frontend geminizer-frontend=$(DOCKER_REGISTRY)/geminizer-frontend:$(VERSION)

# Platform Specific
wsl2: install
	@echo "WSL2 setup complete - run 'make dev' to start"

windows: install
	@echo "Windows setup complete - run 'make dev' to start"

linux: install
	@echo "Linux setup complete - run 'make dev' to start"

termux:
	@echo "Installing on Termux (Android)..."
	pkg install golang nodejs
	$(MAKE) install
	@echo "✅ Termux installation complete"

# Security & Auditing
security-scan:
	@echo "Running comprehensive security scan..."
	trivy image $(DOCKER_REGISTRY)/geminizer-backend:$(VERSION)
	trivy image $(DOCKER_REGISTRY)/geminizer-frontend:$(VERSION)
	gosec -exclude-dir=frontend ./backend/...
	npm audit --production

audit:
	@echo "Running full security audit..."
	$(MAKE) test-security
	$(MAKE) security-scan
	@echo "✅ Security audit complete"

# Maintenance
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/ dist/ node_modules/ backend/vendor/
	docker system prune -f

update-deps:
	@echo "Updating dependencies..."
	cd backend && go get -u all && go mod tidy
	cd frontend && npm update
	$(MAKE) audit

# Help
help:
	@echo "Geminizer Enterprise Build System"
	@echo ""
	@echo "Targets:"
	@echo "  install      - Install all dependencies"
	@echo "  dev          - Start development environment"
	@echo "  build        - Build production binaries"
	@echo "  test         - Run all tests"
	@echo "  deploy       - Deploy to production"
	@echo "  security-scan - Run security vulnerability scan"
	@echo "  audit        - Full security audit"
	@echo "  clean        - Clean build artifacts"
	@echo ""
	@echo "Platforms:"
	@echo "  wsl2         - Setup for WSL2"
	@echo "  windows      - Setup for Windows"
	@echo "  linux        - Setup for Linux"
	@echo "  termux       - Setup for Android Termux"

.DEFAULT_GOAL := help
