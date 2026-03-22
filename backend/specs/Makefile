.PHONY: help

# Global variables
GCP_PROJECT=ekolivs

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

############
# generate #
############

all-protogen: demand-protogen product-protogen order-protogen ## Generate go code for all specs from protobuf specifications

demand-protogen: ## Generate go code for demand specs from protobuf specifications
	@protoc --proto_path=../../specs \
		--go_out=demand_api --go_opt=paths=source_relative \
		--go-grpc_out=demand_api --go-grpc_opt=paths=source_relative \
		demand.proto
	@# Inject custom tags
	@protoc-go-inject-tag -input=demand_api/demand.pb.go
	@echo "Generated protogen for demand_api"

product-protogen: ## Generate go code for product specs from protobuf specifications
	@protoc --proto_path=../../specs \
		--go_out=product_api --go_opt=paths=source_relative \
		--go-grpc_out=product_api --go-grpc_opt=paths=source_relative \
		product.proto
	@protoc-go-inject-tag -input=product_api/product.pb.go
	@echo "Generated protogen for product_api"

order-protogen: ## Generate go code for order specs form protobuf specifications
	@protoc --proto_path=../../specs \
		--go_out=order_api --go_opt=paths=source_relative \
		--go-grpc_out=order_api --go-grpc_opt=paths=source_relative \
		order.proto
	@protoc-go-inject-tag -input=order_api/order.pb.go
	@echo "Generated protogen for order_api"
