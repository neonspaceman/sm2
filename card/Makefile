.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

include ./.env
MIGRATIONS_FOLDER=./migrations

.PHONY: help
help: ## Show commands descriptions
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(abspath $(firstword $(MAKEFILE_LIST))) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-38s\033[0m %s\n", $$1, $$2}'

.PHONY: database-create
database-create: ## Create database
	PGPASSWORD=$(POSTGRES_PASSWORD) docker compose run --rm postgresql psql -h postgresql -U $(POSTGRES_USER) -tc "SELECT 1 FROM pg_database WHERE datname='$(POSTGRES_DATABASE)'" | grep -q 1 \
		|| PGPASSWORD=$(POSTGRES_PASSWORD) docker compose run --rm postgresql -h postgresql psql -U $(POSTGRES_USER) -c "CREATE DATABASE $(POSTGRES_DATABASE)"

.PHONY: migrate-up
migrate-up: ## Execute migrations
	docker compose run --rm --volume ./:/app card \
		/go/bin/goose -dir $(MIGRATIONS_FOLDER) postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" up

.PHONY: migrate-down
migrate-down: ## Rollback migration
	docker compose run --rm card \
		/go/bin/goose -dir $(MIGRATIONS_FOLDER) postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" down

.PHONY: migrate-create
migrate-create: ## Create new migration. Example: make migrate-create name=<NAME>
	docker compose run --rm --user $$(id -u):$$(id -g) --volume ./:/app card \
		/go/bin/goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

.PHONY: proto-generate
proto-generate: ## Generate all proto files
	docker compose run --rm --user $$(id -u):$$(id -g) --volume ./:/app card \
		protoc \
			--proto_path=proto \
			--plugin=protoc-gen-go=/go/bin/protoc-gen-go \
         	--go_out=pkg/api --go_opt=paths=source_relative \
         	--plugin=protoc-gen-go-grpc=/go/bin/protoc-gen-go-grpc \
         	--go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
         	card/messages.proto \
         	card/service.proto
