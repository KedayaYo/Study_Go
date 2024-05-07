package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1、grpc 连接
	//grpc.Dial("localhost:8020", grpc.WithInsecure())// 已弃用
	conn, err := grpc.Dial("localhost:8020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer conn.Close()
	// 2、实例化客户端
	client := proto.NewHelloServiceClient(conn)
	response, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name:  "Kedaya",
		Age:   18,
		Hobby: []string{"Game", "Drink"},
	})
	if err != nil {
		panic("failed to call: " + err.Error())
	}
	fmt.Println(response.Reply)
}
