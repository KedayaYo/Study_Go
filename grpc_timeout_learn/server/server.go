package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_timeout_learn/proto/pg_go"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (rsp *proto.HelloResponse, err error) {
	// 模拟超时
	time.Sleep(time.Second * 5)
	reply := fmt.Sprintf("Hello %s, I'm %d years old, I like %v", req.Name, req.Age, req.Hobby)
	return &proto.HelloResponse{Reply: reply}, nil
}
func main() {
	//1、实例化grpc Server
	g := grpc.NewServer()
	proto.RegisterHelloServiceServer(g, &Server{})
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
