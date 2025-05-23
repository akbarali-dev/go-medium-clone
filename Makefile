# Makefile

# Run go mod tidy to clean up dependencies

-include .env

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

tidy:
	@go mod tidy
	@go mod vendor
run:
	@go run cmd/main.go

migration:
	@migrate create -ext sql -dir ./migrations -seq $(name)

migrateup:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose up

migratedown:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose down
test:
	@go test ./... -v
test-cover:
	@go test ./... -cover -v