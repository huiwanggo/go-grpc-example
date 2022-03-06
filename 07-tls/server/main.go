package main

import (
	"context"
	"fmt"
	pb "github.com/huiwanggo/go-grpc-example/07-tls/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

// SimpleService 定义服务
type SimpleService struct {
	pb.UnimplementedSimpleServer
}

// Send 实现服务方法
func (s *SimpleService) Send(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	fmt.Println("接收到请求：" + in.Data)
	res := pb.SimpleResponse{
		Code:  200,
		Value: "测试：" + in.Data,
	}
	return &res, nil
}

func main() {
	// 监听本地端口
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	// tls
	creds, err := credentials.NewServerTLSFromFile("tls/cert.pem", "tls/key.pem")
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}
	// 创建gRPC服务器实例
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 注册服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	// 启动服务 阻塞等待请求
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
