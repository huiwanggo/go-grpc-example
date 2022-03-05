package main

import (
	"context"
	pb "github.com/huiwanggo/go-grpc-example/03-server-stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 建立grpc连接
	grpcClient := pb.NewStreamServerClient(conn)
	req := pb.SimpleRequest{
		Data: "server stream",
	}

	// 调用服务方法
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 接收服务端消息
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.String())
	}

	// CloseSend 关闭服务端的stream，让它停止发送数据, 若继续调用Recv()，会重新激活stream
	//err = stream.CloseSend()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
