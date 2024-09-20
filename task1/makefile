DB_SN := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_SN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-version:
	$(MIGRATE) version

force: 
	$(MIGRATE) force ${VERSION}

run:
	go run cmd/app/main.go 

lint:
	golangci-lint run --out-format=colored-line-number