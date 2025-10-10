# Variables
APP_NAME := main
CMD_PATH := ./cmd/app/main.go
BUILD_DIR := ./build/bin
DOCKER_DIR := ./deploy
WIRE_PATH := ./...
DB_URL = root:nhbookroot@tcp(localhost:3306)/nhbook?charset=utf8mb4&parseTime=True&loc=Local
MIGRATION_DIR = db/migrations

.PHONY: migrate-up migrate-down migrate-status migrate-redo migrate-create

# Chạy tất cả migration chưa chạy
migrate-up:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" up

# Rollback 1 migration
migrate-down:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" down

# Xem trạng thái migration
migrate-status:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" status

# Rollback rồi chạy lại migration gần nhất
migrate-redo:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" redo

migrate-create:
ifndef name
	$(error name is not set. Usage: make migrate-create name=create_orders_table)
endif
	goose -dir $(MIGRATION_DIR) create $(name) sql

# Go Commands
build:
	@echo "Building Go binary..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

run:generate-docs
	@echo "Running Go application..."
	@go run $(CMD_PATH)

wire:
	@echo "Running wire injection..."
	@wire $(WIRE_PATH)

# Docker Commands
docker-up:
	@echo "Starting Docker containers..."
	@cd $(DOCKER_DIR) && docker compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	@cd $(DOCKER_DIR) && docker compose down

docker-restart:
	@echo "Restarting Docker containers..."
	@cd $(DOCKER_DIR) && docker compose restart

generate-docs:
	@echo "Generating API documentation..."
	@swag init -g ./cmd/app/main.go
# Aliases
rundb: docker-up