package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_learn/proto/pg_go"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

// 严格按照生成的proto go文件里面的接口实现 RegisterHelloServiceServer里面的HelloServiceServer
func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (rsp *proto.HelloResponse, err error) {
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
