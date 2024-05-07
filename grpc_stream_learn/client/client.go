package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_stream_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

// 端口
const HOST = "localhost"
const PORT = ":8020"

// TODO 客户端整体流程
// 1、grpc 连接 获取连接
// 2、proto.NewGreeterClient实例化客户端 获取client
// 3、通过client调用对应方法 获取stream
// 4、接收：stream.Recv()接收返回对应的结构体，发送：stream.Send()发送对应的结构体
func main() {
	// 1、grpc 连接
	conn, err := grpc.Dial(HOST+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	// 关闭连接
	defer conn.Close()
	// 2、实例化客户端
	client := proto.NewGreeterClient(conn)
	// TODO 服务端流模式
	// 3、调用  不是根据 type GreeterClient interface里面的方法 GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) 来的
	// 是根据 proto 文件里面的服务方法来的rpc GetStream(StreamReqData)
	getStream, err := client.GetStream(context.Background(), &proto.StreamReqData{
		Data: "客户端请求数据",
	})
	if err != nil {
		panic("failed to call GetStream: " + err.Error())
	}
	// 4、接收数据
	for {
		reply, err := getStream.Recv()
		if err != nil {
			break
		}
		// 打印数据
		fmt.Printf("接收到的数据：%v\n", reply.Reply)
	}
	// TODO 客户端流模式
	// 3、调用
	putStream, err := client.PutStream(context.Background())
	if err != nil {
		panic("failed to call PutStream: " + err.Error())
	}
	// 4、发送数据
	i := 0
	for {
		if i++; i > 10 {
			fmt.Println("客户端发送数据结束")
			break
		}
		err := putStream.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("%v---客户端第%d次发送", time.Now().Format("2006年01月02日 15:04:05"), i),
		})
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	// TODO 双向流模式
	// 3、调用
	allStream, err := client.AllStream(context.Background())
	if err != nil {
		panic("failed to call AllStream: " + err.Error())
	}
	// 4、发送数据
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		// 客户端接收数据
		for {
			recv, err := allStream.Recv()
			if err != nil {
				fmt.Println("服务端接收数据失败：", err)
				break
			}
			fmt.Printf("接收到的数据：%v\n", recv.Reply)
		}
	}()
	go func() {
		// 客户端发送数据
		defer wg.Done()
		i := 0
		for {
			if i++; i > 10 {
				fmt.Println("客户端发送数据结束")
				break
			}
			_ = allStream.Send(&proto.StreamReqData{
				Data: fmt.Sprintf("%v---客户端第%d次发送", time.Now().Format("2006年01月02日 15:04:05"), i),
			})
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
