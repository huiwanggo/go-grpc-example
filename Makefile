pb01:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 01-proto/proto/greeter.proto

pb02:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 02-simple/proto/simple.proto

pb03:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 03-server-stream/proto/server_stream.proto