package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_metadata/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 1、grpc 连接
	conn, err := grpc.Dial("localhost:8020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer conn.Close()
	// 2、实例化客户端
	client := proto.NewHelloServiceClient(conn)
	// 生成 metadata 数据
	//var timestampFormat = time.DateTime
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	m := map[string]string{
		"name":     "Kedaya",
		"password": "123456",
	}
	mslice := make([]string, 0, len(m))
	for k, v := range m {
		mslice = append(mslice, k, v)
	}
	fmt.Printf("mslice: %v\n", mslice)
	//ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 推荐使用AppendToOutgoingContext
	ctx := metadata.AppendToOutgoingContext(context.Background(), mslice...)
	// 将含有metadata的ctx传入
	response, err := client.SayHello(ctx, &proto.HelloRequest{
		Name:  "Kedaya",
		Age:   18,
		Hobby: []string{"Game", "Drink"},
	})
	if err != nil {
		panic("failed to call: " + err.Error())
	}
	fmt.Println(response.Reply)
}
