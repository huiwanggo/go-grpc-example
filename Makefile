pb01:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 01-proto/proto/greeter.proto