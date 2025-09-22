.PHONY: help

# Global variables
GCP_PROJECT=ekolivs

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

############
# generate #
############

demand-protogen: ## Generate go code for demand specs from protobuf specifications
	protoc --proto_path=../../specs \
		--go_out=demand_api --go_opt=paths=source_relative \
		--go-grpc_out=demand_api --go-grpc_opt=paths=source_relative \
		demand.proto
	# Inject custom tags
	protoc-go-inject-tag -input=demand_api/demand.pb.go
