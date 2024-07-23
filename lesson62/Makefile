CURRENT_DIR := $(shell pwd)
DATABASE_URL="postgres://postgres:root@localhost:5432/lesson62?sslmode=disable"

run:
	@go run cmd/main.go

tidy:
	@go mod tidy
	@go mod vendor

mig-create:
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
	fi; \
	migrate create -ext sql -dir db/migrations -seq $$name

mig-up:
	@migrate -database "$(DATABASE_URL)" -path db/migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path db/migrations down

mig-force:
	@if [ -z "$(version)" ]; then \
		read -p "Enter migration version: " version; \
	fi; \
	migrate -database "$(DATABASE_URL)" -path db/migrations force $$version