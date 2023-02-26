BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

.DEFAULT_GOAL := help

.PHONY: compile clean

compile: ## creating pb file and build golang
	@protoc -I greet/${PROTO_DIR} --go_out=. --go_opt=module=${PACKAGE} --go-grpc_out=. --go-grpc_opt=module=${PACKAGE} greet/${PROTO_DIR}/*.proto

clean: ## clean all pb file
	@find . -type f -name '*.pb.go' -exec rm {} +

help: ## show help
	@${HELP_CMD}