package main

import (
	pb "go-grpc-example/03-server-stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

// StreamService 定义服务
type StreamService struct {
	pb.UnimplementedStreamServerServer
}

// ListValue 实现服务方法
func (s *StreamService) ListValue(in *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for i := 0; i < 5; i++ {
		// 向流中发送数据
		err := srv.Send(&pb.StreamResponse{
			Value: in.Data + " : " + strconv.Itoa(i),
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	// 监听本地端口
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	// 创建 grpc 服务实例
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})
	// 启动服务
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
