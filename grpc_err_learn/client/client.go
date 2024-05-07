package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_err_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	response, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name:  "Kedaya",
		Age:   18,
		Hobby: []string{"Game", "Drink"},
	})
	if err != nil {
		// 最简单的错误处理
		//panic("failed to call: " + err.Error())
		// 这里是将err转换为status对象，然后通过status对象获取错误信息
		st, ok := status.FromError(err)
		if !ok {
			panic("解析错误失败")
		}
		fmt.Printf("ErrorCode: %v, ErrorMessage: %v\n", st.Code(), st.Message())
		return
	}
	fmt.Println(response.Reply)
}
