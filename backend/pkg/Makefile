.PHONY: help zettle

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
zettle: ## Generate go code from zettle openapi specification
	oapi-codegen -package zettle -generate types -include-operation-ids getAllProductsV2 https://products.izettle.com/openapi.json > ./zettle/zettle.gen.go