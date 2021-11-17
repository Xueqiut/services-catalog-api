conn_str = postgres://postgres:password@database:5432/postgres?sslmode=disable

migrate:
	docker compose down -v --remove-orphans; docker compose up -d

run:
	go run cmd/server/main.go

seeddb:
	make migrate;
	docker exec -it services-catalog-api-postgres psql "$(conn_str)" -f /testdata/testdata.sql

test:
	make migrate; (cd cmd/server; go test . -count=1)
	
all: seeddb run