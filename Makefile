conn_str = postgres://postgres:password@database:5432/postgres?sslmode=disable

migrate:
	docker compose down -v --remove-orphans; docker compose up -d

seeddb:
	make migrate;
	docker exec -it services-catalog-api-postgres psql "$(conn_str)" -f /testdata/testdata.sql

test:
	make migrate; (cd cmd/server; go test . -count=1)

run:
	make seeddb; go run cmd/server/main.go