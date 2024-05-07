package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_timeout_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
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
	// 3、设置超时机制 使用context
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	// 4、调用方法
	response, err := client.SayHello(ctx, &proto.HelloRequest{
		Name:  "Kedaya",
		Age:   18,
		Hobby: []string{"Game", "Drink"},
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析错误失败")
		}
		fmt.Printf("ErrorCode: %v, ErrorMessage: %v\n", st.Code(), st.Message())
		return
	}
	fmt.Println(response.Reply)
}
