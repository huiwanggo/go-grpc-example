package main

import (
	"context"
	"fmt"
	pb "github.com/huiwanggo/go-grpc-example/08-token/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// Token 定义 PerRPCCredentials
type Token struct {
	APPID  string
	Secret string
	Token  string
}

func (t Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_id": t.APPID, "secret": t.Secret, "token": t.Token}, nil
}

func (t Token) RequireTransportSecurity() bool {
	return true
}

// LogInterceptor 客户端拦截器
func LogInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("call " + method)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func main() {

	T := Token{
		APPID:  "test-appid",
		Secret: "test-secret",
	}
	T.Token = "test-token"

	// tls
	creds, err := credentials.NewClientTLSFromFile("tls/cert.pem", "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// 连接服务器 添加 WithPerRPCCredentials Option， 拦截器 WithUnaryInterceptor
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(T), grpc.WithUnaryInterceptor(LogInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 建立 grpc 连接
	grpcClient := pb.NewSimpleClient(conn)

	// 请求结构体
	req := pb.SimpleRequest{
		Data: "xxx",
	}
	// 调用服务方法
	res, err := grpcClient.Send(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	// 打印返回值
	log.Println(res)
}
