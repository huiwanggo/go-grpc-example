package main

import (
	"context"
	pb "github.com/huiwanggo/go-grpc-example/06-timeout/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	// 调用服务方法 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := grpcClient.Send(ctx, &req)

	if err != nil {
		// 获取错误状态
		errStatus, ok := status.FromError(err)
		if ok {
			// 判断调用是否超时
			if errStatus.Code() == codes.DeadlineExceeded {
				log.Fatal("timeout")
			}
		}

		log.Fatal(err)
	}
	// 打印返回值
	log.Println(res)
}
