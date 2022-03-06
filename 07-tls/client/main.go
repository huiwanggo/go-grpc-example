package main

import (
	"context"
	pb "github.com/huiwanggo/go-grpc-example/07-tls/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	// tls
	creds, err := credentials.NewClientTLSFromFile("tls/cert.pem", "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// 连接服务器
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(creds))
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
