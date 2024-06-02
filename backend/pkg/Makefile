.PHONY: help zettle

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
zettle-product: ## Generate go code from zettle product library openapi specification
	@cd ./zettle && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 --config=config.yaml https://products.izettle.com/openapi.json
	@./scripts/product-env.sh

zettle-inventory: ## Generate go code from zettle inventory openapi specification
	@go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -package zettle -generate types https://inventory.izettle.com/openapi.json > ./zettle/zettle-inventory.gen.go
