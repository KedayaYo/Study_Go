package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_protos_learn/proto/pg_go"
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
	// 调用空的内置empty
	//pong, _ := client.Ping(context.Background(),&empty.Empty{})
	// 嵌套的结构体方法1:
	//proto.Result_Version{}
	// 枚举
	userServeiceClient := proto.NewUserServeiceClient(conn)
	result, errors := userServeiceClient.Introduce(context.Background(), &proto.UserInfo{
		Id:     "1",
		Name:   "跳跳虎",
		Age:    23,
		Gender: proto.Gender_MALE,
		Hobby:  map[string]string{"1": "Game", "2": "Drink", "3": "Code"},
	})
	if errors != nil {
		fmt.Printf("errors:%v\n", errors)
	}
	fmt.Printf("result:%v\n", result)

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
