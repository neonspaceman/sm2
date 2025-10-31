.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

ENV_FILE=.env
include $(ENV_FILE)

.PHONY: help
help: ## Show commands descriptions
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(abspath $(firstword $(MAKEFILE_LIST))) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-38s\033[0m %s\n", $$1, $$2}'

.PHONY: database-create
database-create: ## Create database
	cd ../infrastructure \
		&& PGPASSWORD=$(POSTGRES_PASSWORD) docker exec $$(docker compose ps -q postgresql) psql -U $(POSTGRES_USER) -tc "SELECT 1 FROM pg_database WHERE datname='$(POSTGRES_DATABASE)'" | grep -q 1 \
		|| PGPASSWORD=$(POSTGRES_PASSWORD) docker exec $$(docker compose ps -q postgresql) psql -U $(POSTGRES_USER) -c "CREATE DATABASE $(POSTGRES_DATABASE)"

.PHONY: migrate-up
migrate-up: ## Execute migrations
	docker run \
		--rm \
		--volume ./:/app \
		--network $(DOCKER_NETWORK_NAME) \
		$$(docker compose images -q telegram-bot) \
		/go/bin/goose \
			-dir $(MIGRATIONS) postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" up

.PHONY: migrate-down
migrate-down: ## Rollback migration
	docker run \
		--rm \
		--volume ./:/app \
		--network $(DOCKER_NETWORK_NAME) \
		$$(docker compose images -q telegram-bot) \
		/go/bin/goose \
 			-dir $(MIGRATIONS) postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DATABASE}?sslmode=disable" down

.PHONY: migrate-create
migrate-create: ## Create new migration. Example: make migrate-create name=<NAME>
	docker run \
		--rm \
		--user $$(id -u):$$(id -g) \
		--volume ./:/app \
		$$(docker compose images -q telegram-bot) \
		/go/bin/goose \
	  		-dir $(MIGRATIONS) create $(name) sql
