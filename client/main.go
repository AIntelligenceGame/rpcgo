package main

import (
	"context"
	"log"

	pb "github.com/shangguanairen/rpcgo/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 建立连接，没有加密验证
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 创建客户端
	client := pb.NewSayHelloClient(conn)
	// 远程调用
	helloRep, err := client.Hello(context.Background(), &pb.HelloReq{Name: "client"})
	if err != nil {
		panic(err)
	}
	log.Printf("received grpc resp: %+v", helloRep.String())
}
