build:
	@go build -o bin/mdpages ./cmd/mdpages/main.go
	@cp .env bin

run: build
	@DEBUG=true ./bin/mdpages

test:
	@go test

lint:
	@golangci-lint run

docker:
	@docker build -t yosaa5782/mdpages .
	@docker run --rm --name mdpages-web -p 5000:5000 --network mdpages-net -d yosaa5782/mdpages

postgres:
	@docker run --rm \
		--name mdpages-postgres \
		-p 5432:5432 \
		-e POSTGRES_USER=user \
		-e POSTGRES_PASSWORD=1234 \
		-v postgres-volume:/var/lib/postgresql/data \
		--network mdpages-net \
		-d postgres
	@docker run --rm \
		--name mdpages-adminer \
		-p 5050:8080 \
		--network mdpages-net \
		-d adminer

create-db:
	@docker exec -it mdpages-postgres createdb --username=user --owner=user mdpages_db

create-network:
	@docker network create -d bridge mdpages-net

mig:
	@migrate create -ext sql -dir migrations -seq migration_name

migrate-up:
	@migrate -path migrations -database "postgres://user:1234@localhost:5432/mdpages_db?sslmode=disable" -verbose up

migrate-down:
	@migrate -path migrations -database "postgres://user:1234@localhost:5432/mdpages_db?sslmode=disable" -verbose down

migrate-fix:
	@migrate -path migrations -database "postgres://user:1234@localhost:5432/mdpages_db?sslmode=disable" force VERSION

redis:
	@docker run --rm \
		--name mdpages-redis \
		-p 6379:6379 \
		-d redis

redis-cli:
	@docker exec -it mdpages-redis redis-cli
