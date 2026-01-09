.PHONY: help build up down restart logs clean dev-vue dev-gin install test lint

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Docker commands
build: ## Build all Docker images
	docker-compose build

up: ## Start all services
	docker-compose up -d

down: ## Stop all services
	docker-compose down

restart: ## Restart all services
	docker-compose restart

logs: ## Show logs from all services
	docker-compose logs -f

logs-gin: ## Show logs from Gin service only
	docker-compose logs -f gin

logs-vue: ## Show logs from Vue service only
	docker-compose logs -f vue

logs-mongo: ## Show logs from MongoDB service only
	docker-compose logs -f mongo

logs-redis: ## Show logs from Redis service only
	docker-compose logs -f redis

# Development commands
dev: ## Start the full application (build + up)
	docker-compose up --build

dev-detached: ## Start the full application in background
	docker-compose up --build -d

dev-vue: ## Start Vue development server locally
	cd vue && npm run dev

dev-gin: ## Start Gin development server locally
	cd gin && go run main.go

# Installation commands
install: ## Install dependencies for both frontend and backend
	cd vue && npm install
	cd gin && go mod download

install-vue: ## Install Vue dependencies only
	cd vue && npm install

install-gin: ## Install Go dependencies only
	cd gin && go mod download

# Build commands
build-vue: ## Build Vue application for production
	cd vue && npm run build

build-gin: ## Build Go application
	cd gin && go build -o main .

# Testing commands
test: ## Run tests for both frontend and backend
	cd vue && npm run lint
	cd gin && go test ./...

test-vue: ## Run Vue tests and linting
	cd vue && npm run lint

test-gin: ## Run Go tests
	cd gin && go test ./...

# Utility commands
clean: ## Clean up Docker containers, images, and volumes
	docker-compose down -v
	docker system prune -f

clean-all: ## Clean up everything including unused images
	docker-compose down -v
	docker system prune -af

status: ## Show status of all services
	docker-compose ps

shell-gin: ## Open shell in Gin container
	docker-compose exec gin sh

shell-vue: ## Open shell in Vue container
	docker-compose exec vue sh

shell-mongo: ## Open MongoDB shell
	docker-compose exec mongo mongosh

shell-redis: ## Open Redis CLI
	docker-compose exec redis redis-cli

# Database commands
db-reset: ## Reset database (stop, remove volumes, start)
	docker-compose stop mongo redis
	docker-compose rm -f mongo redis
	docker volume prune -f
	docker-compose up -d mongo redis

# Production commands
prod: ## Start production environment
	docker-compose -f docker-compose.yml up --build -d

prod-logs: ## Show production logs
	docker-compose logs -f

# Quick commands
start: up ## Alias for up
stop: down ## Alias for down
rebuild: clean dev ## Clean and rebuild everything