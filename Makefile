migrate:
	docker compose down; docker compose up -d

build:
	go build cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	docker compose down; docker compose up -d; (cd cmd/server; go test .)
	
all: migrate run