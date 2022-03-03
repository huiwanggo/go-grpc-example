package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/04-client-stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func main() {
	// 连接服务
	conn, _ := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	// 建立grpc连接
	streamClient := pb.NewStreamClientClient(conn)
	// 调用服务方法
	stream, _ := streamClient.List(context.Background())
	for i := 0; i < 5; i++ {
		// 向流中发送数据
		_ = stream.Send(&pb.StreamRequest{Data: "stream client rpc " + strconv.Itoa(i)})
	}
	// 关闭流并获取返回的数据
	res, _ := stream.CloseAndRecv()
	// 打印
	fmt.Println(res)
}
