package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/05-stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"strconv"
)

func main() {
	// 建立连接
	conn, _ := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	// 连接grpc
	streamClient := pb.NewStreamClient(conn)
	// 调用服务方法获得流
	stream, _ := streamClient.Message(context.Background())

	for i := 0; i < 5; i++ {
		// 发送流
		err := stream.Send(&pb.StreamRequest{Value: "client: " + strconv.Itoa(i)})
		if err != nil {
			log.Fatal(err)
		}

		// 接收流
		res, err := stream.Recv()
		// 判断是否流已关闭
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		// 打印结果
		fmt.Println(res.Value)
	}

	// 关闭流
	_ = stream.CloseSend()

}
