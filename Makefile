.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

#CURRENT_USER := $(shell id -u)
ENV_FILE=$(PWD)/.env
include $(ENV_FILE)

.PHONY: help
help: ## Show commands descriptions
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(abspath $(firstword $(MAKEFILE_LIST))) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-38s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build docker image
	docker compose build

.PHONY: up
up: ## Run go app
	docker compose up --force-recreate

.PHONY: down
down: ## Down all infrastructure containers
	docker compose down

# DEV infrastructure enviroment

.PHONY: infrastructure-create-network
infrastructure-create-network: ## Create network
	$(eval NETWORK=$(shell docker network ls | grep $(DOCKER_NETWORK_NAME) | wc -l))
	@if [ $(NETWORK) -eq 0 ]; then \
  		docker network create \
			--attachable \
			--subnet="$(DOCKER_NETWORK_CIDR)" \
			--driver=bridge \
			-o com.docker.network.bridge.name="$(DOCKER_NETWORK_BRIDGE_NAME)" \
			-o com.docker.network.driver.mtu="$(DOCKER_NETWORK_MTU)" \
			$(DOCKER_NETWORK_NAME) ; \
	fi

.PHONY: infrastructure-up
infrastructure-up: ## Up all infrastructure containers
	make infrastructure-create-network && \
	docker compose -f ./infrastructure/docker-compose.yml up -d

.PHONY: infrastructure-down
infrastructure-down: ## Down all infrastructure containers
	docker compose -f ./infrastructure/docker-compose.yml down
