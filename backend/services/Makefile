# Extract the first command-line goal as the actual target
COMMAND := $(firstword $(MAKECMDGOALS))
TARGET := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

INITIAL_PORT := 8080
COMMAND_SUFFIX=main.go
ifdef f
	COMMAND_SUFFIX=main.go -f $(f)
endif


.PHONY: help rund $(TARGET)

# Global variables
GCP_PROJECT = ekolivs

# Expand now; ALL_SERVICES is now the result of $(shell ...)
ALL_SERVICES := $(shell ls -d */ | tr -d "/")
help: ## list all the Makefile commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###########
# general #
###########
run: setgcpenv ## Run  service with go directly
	@echo "Use the local makefile to instantiate the service one by one"

rund: ## Run service with passing absolute path of a filename as an argument: make rund $service_name f="/absolute_path/filename"
	@echo "$(TARGET) to be run: "
ifeq (,$(shell which ss))
	$(error "Error: 'ss' command not found in PATH. Please install it.")
endif

ifeq "$(findstring $(TARGET), $(ALL_SERVICES))" ""
	@echo "Service is not yet defined"
	@exit 1
endif

$(TARGET): rund setgcpenv
	@if [ ! -f $(TARGET)/cmd/main.go ]; then \
		echo "Service entrypoint $(TARGET)/cmd/main.go is not defined"; \
		exit 1; \
	fi
	@POTENTIAL_PORT=$(INITIAL_PORT); \
	FOUND=1; \
	while [ $$FOUND -eq 1 ]; do \
	   if ss -tlnu | grep -q ":$${POTENTIAL_PORT}"; then \
	      POTENTIAL_PORT=$$(expr $$POTENTIAL_PORT + 1); \
	   else \
	      FOUND=0; \
	   fi; \
	done; \
	export PORT=$${POTENTIAL_PORT}; \
	echo "$(TARGET) is going to be bounded to $${PORT}"; \
	go run $(TARGET)/cmd/$(COMMAND_SUFFIX) &
	

protogen: ## Generate go code from protobuf specifications
	cd ../../specs && make demand-protogen

#######
# GCP #
#######

setgcpenv:
	$(eval export GOOGLE_CLOUD_PROJECT=${GCP_PROJECT})

# This rule makes any extra "arguments" on the command line effectively do-nothing targets,
# preventing 'make' from trying to build them as actual files.
%:: @true
