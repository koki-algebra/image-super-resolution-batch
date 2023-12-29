.PHONY: help
.DEFAULT_GOAL := help

.PHONY: build
build: ## build application
	@docker compose build

run: ## start application
	@docker compose up -d

generate: ## generate code
	@cd gateway/api/http && oapi-codegen -config config.yml openapi.yml

fmt: ## format code
	@cd gateway && go fmt ./...

clear: ## clear application
	@docker compose down --volumes

logs: ## show API server logs
	@docker compose logs -f gateway

help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
