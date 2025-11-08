.PHONY: help

# Global variables
GCP_PROJECT=ekolivs

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
run: setgcpenv ## Run demand service with go directly
	go run ./cmd/main.go

rund: setgcpenv ## Run demand service with passing absolute path of a filename as an argument: make runf f="/absolute_path/filename"
ifdef f
	go run ./cmd/main.go -f $(f)
else
	@echo "Filename and its absolute path are required"
	@exit 1
endif

protogen: ## Generate go code from protobuf specifications
	cd ../../specs && make demand-protogen

#######
# GCP #
#######

setgcpenv:
	$(eval export GOOGLE_CLOUD_PROJECT=${GCP_PROJECT})

