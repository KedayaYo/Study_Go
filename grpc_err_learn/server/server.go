package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_err_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (rsp *proto.HelloResponse, err error) {
	//reply := fmt.Sprintf("Hello %s, I'm %d years old, I like %v", req.Name, req.Age, req.Hobby)
	//return &proto.HelloResponse{Reply: reply}, nil
	// status.Errorf底层使用的就是status.New  一般报错使用status.Error或者status.Errorf即可
	return nil, status.Errorf(codes.NotFound, "记录未找到: %s", req.Name)
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
