.PHONY: help zettle

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
zettle-product-library: ## Generate go code from zettle product library openapi specification
	oapi-codegen -package zettle -generate types -include-operation-ids getAllProductsV2,updateProduct https://products.izettle.com/openapi.json > ./zettle/zettle-product.gen.go

zettle-inventory: ## Generate go code from zettle inventory openapi specification
	oapi-codegen -package zettle -generate types https://inventory.izettle.com/openapi.json > ./zettle/zettle-inventory.gen.go
