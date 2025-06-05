.PHONY: help

# Global variables
GCP_PROJECT=ekolivs

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
run: setgcpenv ## Run order service with go directly
	go run ./cmd/main.go

protogen: ## Generate go code from protobuf specifications
	protoc --proto_path=../../../specs \
		--go_out=order_api --go_opt=paths=source_relative \
		--go-grpc_out=order_api --go-grpc_opt=paths=source_relative \
		order.proto

#######
# GCP #
#######

setgcpenv:
	$(eval export GOOGLE_CLOUD_PROJECT=${GCP_PROJECT})
