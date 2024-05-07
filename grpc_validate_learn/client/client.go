package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_validate_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1、创建intercept
	helloUnaryIntercept := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("客户端拦截到了一个请求%s\n", method)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	// 配置集合
	opts := make([]grpc.DialOption, 0, 2)
	// 2、grpc生成客户端连接配置：WithUnaryInterceptor
	opts = append(opts, grpc.WithUnaryInterceptor(helloUnaryIntercept))
	// 3、grpc 连接
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Printf("%v\n", opts)
	conn, err := grpc.Dial("localhost:8020", opts...)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer conn.Close()
	// 4、实例化客户端
	client := proto.NewGreeterClient(conn)
	response, err := client.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "kedaya1011@163.com",
		Mobile: "18888888888",
		Home: &proto.Person_Location{
			Lat: -90.0,
			Lng: 90.0,
		},
	})
	if err != nil {
		panic("failed to SayHello: " + err.Error())
	}
	fmt.Printf("response: %v\n", response)
}
