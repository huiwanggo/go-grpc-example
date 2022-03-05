package main

import (
	"fmt"
	pb "go-grpc-example/05-stream/proto"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
)

type StreamService struct {
	pb.UnimplementedStreamServer
}

func (s *StreamService) Message(srv pb.Stream_MessageServer) error {
	i := 1
	for {
		// 接收请求
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println(req.Value)

		// 发送响应
		err = srv.Send(&pb.StreamResponse{
			Value: "server: " + strconv.Itoa(i) + " " + req.Value,
		})
		if err != nil {
			return err
		}

		i++
	}
}

func main() {
	// 监听端口
	listener, _ := net.Listen("tcp", ":8888")
	// 创建grpc
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterStreamServer(grpcServer, &StreamService{})
	// 启动服务
	_ = grpcServer.Serve(listener)
}
