package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/shangguanairen/rpcgo/hello"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb.UnimplementedSayHelloServer
}

func (g *GrpcServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRep, error) {
	log.Printf("received grpc req: %+v", req.String())
	return &pb.HelloRep{Msg: fmt.Sprintf("hello world! %s", req.Name)}, nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 创建gprc服务器
	server := grpc.NewServer()
	// 注册服务
	pb.RegisterSayHelloServer(server, &GrpcServer{})
	// 运行
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
