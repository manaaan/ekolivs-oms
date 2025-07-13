.PHONY: help

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###################
# setup + tooling #
###################
setup: ## Set up tool dependencies for code generation
	@PROTOC_VERSION=27.0 ./scripts/setup-tools.sh

protogen: ## Generate go code from protobuf specifications
	protoc --proto_path=specs \
		--go_out=backend/services/product/api --go_opt=paths=source_relative \
		--go-grpc_out=backend/services/product/api --go-grpc_opt=paths=source_relative \
		./specs/product.proto
	protoc --proto_path=specs \
		--go_out=backend/services/demand/api --go_opt=paths=source_relative \
		--go-grpc_out=backend/services/demand/api --go-grpc_opt=paths=source_relative \
		./specs/demand.proto
	# Inject custom tags
	protoc-go-inject-tag -input=backend/services/demand/api/demand.pb.go
