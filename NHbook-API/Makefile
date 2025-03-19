rundb:
	@cd deploy && docker compose up -d
build:
	@go build -o ./build/bin/main ./cmd/app/main.go
run:build
	./build/bin/main
wire:
	@go run github.com/google/wire/cmd/wire ./...
