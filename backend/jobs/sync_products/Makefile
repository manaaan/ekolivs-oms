.PHONY: help

help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
run: setgcpenv ## Run product service with go directly
	go run ./cmd/main.go

#######
# GCP #
#######

setgcpenv:
	$(eval export GOOGLE_CLOUD_PROJECT=$(shell gcloud config get-value project;echo $?))
