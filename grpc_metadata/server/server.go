package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_metadata/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (rsp *proto.HelloResponse, err error) {
	// 服务端尝试接收 metadata 数据，通过 FromIncomingContext 接收
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("get metadata error")
	} else {
		fmt.Println(md)
	}
	if name, ok := md["name"]; ok {
		fmt.Printf("name: %v\n", name)
	}
	if password, ok := md["password"]; ok {
		fmt.Printf("password: %v\n", password)
	}
	//for k, v := range md {
	//	fmt.Printf("k: %v, v: %v\n", k, v)
	//}
	reply := fmt.Sprintf("Hello %s, I'm %d years old, I like %v", req.Name, req.Age, req.Hobby)
	return &proto.HelloResponse{Reply: reply}, nil
}
func main() {
	//1、实例化grpc Server
	g := grpc.NewServer()
	proto.RegisterHelloServiceServer(g, &Server{}) // type HelloServiceServer interface是一个接口
	// 启动
	listener, err := net.Listen("tcp", ":8020")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	fmt.Printf("grpc serve start success\n")
	err = g.Serve(listener)
	if err != nil {
		panic("failed to start grpc serve: " + err.Error())
	}
}
