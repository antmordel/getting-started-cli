# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# BUILDING
# ==================================================================================== #

## build: build the service executable
.PHONY: build
build:
	@echo "🚀 Building todocli"
	@go build -o todocli .