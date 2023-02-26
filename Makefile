.PHONY: compile

compile:
	protoc -I greet/proto --go_out=. --go_opt=module=github.com/prolifel/grpc-learn --go-grpc_out=. --go-grpc_opt=module=github.com/prolifel/grpc-learn greet/proto/dummy.proto