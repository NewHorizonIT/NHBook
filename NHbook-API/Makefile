# Variables
APP_NAME := main
CMD_PATH := ./cmd/app/main.go
BUILD_DIR := ./build/bin
DOCKER_DIR := ./deploy
WIRE_PATH := ./...

# Go Commands
build:
	@echo "Building Go binary..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

run:
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

# Aliases
rundb: docker-up