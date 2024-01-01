.PHONY: help
.DEFAULT_GOAL := help

fmt: ## format code
	@rye run black ./src

lint: ## lint code
	@rye run flake8 ./src

help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
