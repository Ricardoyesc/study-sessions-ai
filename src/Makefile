.PHONY: build run test lint dev db-reset docker-up docker-down

APP_NAME = sai-server
CMD_DIR = cmd/server

build:
	go build -o bin/$(APP_NAME) ./$(CMD_DIR)

run:
	go run ./$(CMD_DIR)

test:
	go test ./...

lint:
	golangci-lint run ./...

dev:
	ENVIRONMENT=development go run ./$(CMD_DIR)

db-reset:
	rm -f data/sai_dev.db

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker build -f deploy/docker/Dockerfile -t sai-server:latest .
