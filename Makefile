pb01:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 01-proto/proto/greeter.proto

pb02:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 02-simple/proto/simple.proto

pb03:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 03-server-stream/proto/server_stream.proto

pb04:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 04-client-stream/proto/client_stream.proto

pb05:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 05-stream/proto/stream.proto

pb06:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 06-timeout/proto/timeout.proto

pb07:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 07-tls/proto/tls.proto

pb08:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. 08-token/proto/token.proto

pb:
	make pb01
	make pb02
	make pb03
	make pb04
	make pb05
	make pb06
	make pb07
	make pb08

.PHONY: tls
tls:
	 go run $(GOROOT)/src/crypto/tls/generate_cert.go -host localhost
	 mv cert.pem ./tls/cert.pem
	 mv key.pem ./tls/key.pem
