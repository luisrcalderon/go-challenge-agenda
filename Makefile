.PHONY: proto mock swagger build run-agenda run-api docker-up docker-down test

proto:
	buf generate

mock:
	mockery

swagger:
	cd services/api && $(shell go env GOPATH)/bin/swag init -g cmd/main.go -o docs

build:
	go build ./...

run-agenda:
	go run ./services/agenda/cmd

run-api:
	go run ./services/api/cmd

docker-up:
	docker compose up --build

docker-down:
	docker compose down -v

test:
	go test ./...

