.PHONY: all
all:
	make pb
	make tls

.PHONY: pb
pb:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 01-proto/proto/greeter.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 02-simple/proto/simple.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 03-server-stream/proto/server_stream.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 04-client-stream/proto/client_stream.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 05-stream/proto/stream.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 06-timeout/proto/timeout.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 07-tls/proto/tls.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 08-token/proto/token.proto
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 09-interceptor/proto/token.proto

.PHONY: tls
tls:
	 go run $(GOROOT)/src/crypto/tls/generate_cert.go -host localhost
	 mv cert.pem ./tls/cert.pem
	 mv key.pem ./tls/key.pem
