package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_protos_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"
)

type Server struct {
}

func (s *Server) Introduce(ctx context.Context, info *proto.UserInfo) (*proto.Result, error) {
	msg := fmt.Sprintf("Introduce: %v\n", info)
	fmt.Println(msg)
	// 创建Version实例并初始化
	version := &proto.Result_Version{
		Id:      "1",
		Version: "Version@1.0",
	}

	return &proto.Result{
		Id:         info.Id,
		Code:       proto.Code_OK,
		Version:    version,
		Data:       []string{info.Name, fmt.Sprintf("%d", info.Age), fmt.Sprintf("%v", info.Hobby)},
		Timestamps: timestamppb.New(time.Now()),
	}, nil
}

func (s *Server) Ping(ctx context.Context, empty *emptypb.Empty) (*proto.Pong, error) {
	fmt.Println("Ping")
	return nil, nil
}

// 严格按照生成的proto go文件里面的接口实现 RegisterHelloServiceServer里面的HelloServiceServer
func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (rsp *proto.HelloResponse, err error) {
	reply := fmt.Sprintf("Hello %s, I'm %d years old, I like %v", req.Name, req.Age, req.Hobby)
	return &proto.HelloResponse{Reply: reply}, nil
}
func main() {
	//1、实例化grpc Server
	g := grpc.NewServer()
	// 每个服务都要注册
	proto.RegisterHelloServiceServer(g, &Server{}) // type HelloServiceServer interface是一个接口
	proto.RegisterGoodsServiceServer(g, &Server{}) // type GoodsServiceServer interface是一个接口
	proto.RegisterUserServeiceServer(g, &Server{}) // type UserServeiceServer interface是一个接口
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
