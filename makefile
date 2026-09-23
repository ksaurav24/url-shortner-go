.PHONY: build run build-migrate

build: 
	@go build -o bin/api ./cmd/api

build-migrate:
	@go build -o bin/migrate ./cmd/migration

run: build
	@./bin/api

dev: 
	@air

migrate-up: build-migrate
	@./bin/migrate up

migrate-down: build-migrate
	@./bin/migrate down