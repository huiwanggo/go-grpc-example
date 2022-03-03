package main

import (
	"fmt"
	pb "go-grpc-example/04-client-stream/proto"
	"google.golang.org/grpc"
	"io"
	"net"
)

// SimpleService 定义服务
type SimpleService struct {
	pb.UnimplementedStreamClientServer
}

// List 实现服务方法
func (s *SimpleService) List(srv pb.StreamClient_ListServer) error {
	for {
		// 获得流
		res, err := srv.Recv()
		// 已经完成接收流
		if err == io.EOF {
			// 发送并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}

		if err != nil {
			return err
		}

		fmt.Println(res.String())
	}
}

func main() {
	// 监听端口
	listener, _ := net.Listen("tcp", ":8888")
	// 创建grpc实例
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})
	// 启动服务
	_ = grpcServer.Serve(listener)
}
