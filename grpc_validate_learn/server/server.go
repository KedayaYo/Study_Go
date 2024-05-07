package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_validate_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.Person) (rsp *proto.Person, err error) {

	return &proto.Person{
		Id:     req.Id,
		Email:  req.Email,
		Mobile: req.Mobile,
		Home:   req.Home,
	}, nil
}

// 生成的脚本里已经实现了
type Validator interface {
	Validate() error
}

func main() {

	// 过滤器 简单使用
	//p := new(proto.Person)
	//pErr := p.Validate()
	//if pErr != nil {
	//	panic(pErr)
	//}
	// 拦截器
	helloUnaryInterce := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Printf("服务端拦截到了一个请求%s\n", info.FullMethod)
		// 因为我希望使用Validator来过滤  将req这个any接口转换成req.(Validator)Validator  不能只转换成proto.Person 这样其他接口使用不了
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			} else {
				fmt.Printf("验证通过\n")
			}
		}
		return handler(ctx, req)
	}
	// 2、grpc生成服务端拦截器配置：UnaryInterceptor (stream有对应的拦截器)
	opt := grpc.UnaryInterceptor(helloUnaryInterce)
	// 3、实例化
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
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
