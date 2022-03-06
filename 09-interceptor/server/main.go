package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/huiwanggo/go-grpc-example/08-token/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

	// 获取 Metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no token")
	}
	// 打印数据
	fmt.Println(md)
	fmt.Println(md["token"][0])

	res := pb.SimpleResponse{
		Code:  200,
		Value: "测试：" + in.Data,
	}
	return &res, nil
}

// 四种拦截器
// type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
// type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
// type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
// type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)

// AuthInterceptor 服务端一元拦截器
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("auth intercept")
	// 获取 metadata 数据
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// 鉴权逻辑
		fmt.Println("auth 认证成功")
		fmt.Println(md)
	} else {
		return nil, errors.New("no metadata")
	}

	return handler(ctx, req)
}

// LogInterceptor 打印日志
func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println(info.FullMethod)

	return handler(ctx, req)
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
	// 创建gRPC服务器实例，添加 单一拦截器 UnaryInterceptor Option, 链式拦截器 ChainUnaryInterceptor Option
	grpcServer := grpc.NewServer(grpc.Creds(creds),
		grpc.UnaryInterceptor(AuthInterceptor),
		grpc.ChainUnaryInterceptor(
			AuthInterceptor,
			LogInterceptor,
		))
	// 注册服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	// 启动服务 阻塞等待请求
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
