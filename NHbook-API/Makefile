rundb:
	@cd deploy && docker compose up -d
build:
	@go build -o ./build/bin/main ./cmd/app/main.go
run:
	@go run cmd/app/main.go
wire:
	@go run github.com/google/wire/cmd/wire ./...
