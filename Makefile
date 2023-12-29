.PHONY: help
.DEFAULT_GOAL := help

.PHONY: build
build: ## build application
	@docker compose build --no-cache

run: ## start application
	@docker compose up -d

fmt: ## format code
	@cd gateway && go fmt ./...

clear: ## clear application
	@docker compose down --volumes

logs: ## show API server logs
	@docker compose logs -f gateway

help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
