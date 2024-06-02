.PHONY: help

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###################
# setup + tooling #
###################
setup: ## Set up tool dependencies for code generation
	@PROTOC_VERSION=27.0 ./scripts/setup-tools.sh
