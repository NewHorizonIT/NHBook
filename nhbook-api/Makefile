.PHONY: build run clean lint tidy air migrate-up migrate-down migrate-status migrate-create help

# VARIABLES APP
OUTPUT_DIR := ./bin
APP_NAME := nhbook-api

# VARIABLES GO
GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_TIDY := $(GO) mod tidy
GO_LINT := golangci-lint run

# VARIABLES AIR
AIR := air

# VARIABLES GOOSE
GOOSE := goose
include .env
export

# TARGETS
help:
	@echo "Makefile commands:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application"
	@echo "  clean      - Clean the build artifacts"
	@echo "  lint       - Run code linting"
	@echo "  tidy       - Tidy up Go modules"
	@echo "  air        - Run the application with Air for live reloading"

build:
	@echo "Building the application..."
	$(GO_BUILD) -o $(OUTPUT_DIR)/$(APP_NAME) ./cmd/api/main.go

run:
	@echo "Running the application..."
	$(GO_RUN) ./cmd/api/main.go

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(OUTPUT_DIR)

air: 
	@echo "Running the application with Air..."
	$(AIR) -c .air.toml

# Migration commands
migrate-up: bin/dburl
	@echo "Applying database migrations..."
	@goose -dir ./internal/modules/auth/migrations up
	@goose -dir ./internal/modules/book/migrations up
	@goose -dir ./internal/modules/cart/migrations up
	@goose -dir ./internal/modules/order/migrations up
	@goose -dir ./internal/modules/payment/migrations up

migrate-down:
	@echo "Reverting database migrations..."
	@goose -dir ./internal/modules/auth/migrations down
	@goose -dir ./internal/modules/book/migrations down
	@goose -dir ./internal/modules/cart/migrations down
	@goose -dir ./internal/modules/order/migrations down
	@goose -dir ./internal/modules/payment/migrations down

migrate-status:
	@echo "Checking migration status..."
	goose -dir ./internal/modules/auth/migrations status
	goose -dir ./internal/modules/book/migrations status
migrate-create:
	@echo "Creating a new migration..."
ifndef name
	$(error Migration name is required. Usage: make migrate-create name=<migration_name>)
endif
ifndef module
	$(error Module name is required. Usage: make migrate-create module=<module_name>)
endif
	@echo "Creating migration '$(name)' for module '$(module)'..."
	goose create $(name) sql -dir ./internal/modules/$(module)/migrations

# SQLC generation
sqlc-gen:
ifndef module
	$(error Module name is required. Usage: make sqlc-gen module=<module_name>)
endif
	@echo "Generating SQLC code..."
	sqlc generate -f ./internal/modules/$(module)/sqlc.yaml

# Database seeding
seed:
	@echo "🌱 Seeding database..."
	$(GO_RUN) ./cmd/seeding/seeding.go

seed-build:
	@echo "Building seeding binary..."
	$(GO_BUILD) -o $(OUTPUT_DIR)/seed ./cmd/seeding/seeding.go

# Swagger generation
swagger-gen:
	@echo "Generating Swagger documentation..."
	swag init -g ./cmd/api/main.go -o ./docs --parseDependency --parseInternal