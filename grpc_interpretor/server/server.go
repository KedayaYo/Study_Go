package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_interpretor/proto/pg_go"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (rsp *proto.HelloResponse, err error) {
	reply := fmt.Sprintf("Hello %s, I'm %d years old, I like %v", req.Name, req.Age, req.Hobby)
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {
	// 1、生成参数，参数是一个UnaryServerInterceptor 是一个func 只需要实现这个func即可  handler原本的调用逻辑
	// func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {}
	helloUnaryInterce := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Printf("服务端拦截到了一个请求%s\n", info.FullMethod)
		start := time.Now()
		// 执行初始的逻辑 handler(ctx, req)
		res, err := handler(ctx, req)
		fmt.Printf("服务端耗时：%v\n", time.Since(start))
		return res, err
	}
	// 2、grpc生成服务端拦截器配置：UnaryInterceptor (stream有对应的拦截器)
	opt := grpc.UnaryInterceptor(helloUnaryInterce)
	// 3、实例化
	g := grpc.NewServer(opt)
	proto.RegisterHelloServiceServer(g, &Server{})
	// 4、启动
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
