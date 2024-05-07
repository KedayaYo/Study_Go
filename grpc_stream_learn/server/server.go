package main

import (
	"fmt"
	proto "go_learn/grpc_stream_learn/proto/pg_go"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

// 端口
const PORT = ":8020"

type server struct {
}

func (s server) GetStream(data *proto.StreamReqData, getStream proto.Greeter_GetStreamServer) error {
	// 只要调用 就向客户端发送数据  用for简单演示
	i := 0
	for {
		if i++; i > 10 {
			fmt.Println("服务端发送数据结束")
			break
		}
		_ = getStream.Send(&proto.StreamRspData{
			//Reply: fmt.Sprintf("%v---第%d次发送", time.Now().Unix(), i),
			Reply: fmt.Sprintf("%v---服务端第%d次发送", time.Now().Format("2006年01月02日 15:04:05"), i),
		})
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s server) PutStream(putStream proto.Greeter_PutStreamServer) error {
	// 服务端接收数据
	for {
		recv, err := putStream.Recv()
		if err != nil {
			fmt.Println("服务端接收数据失败：", err)
			break
		}
		fmt.Printf("接收到的数据：%v\n", recv.Data)
	}
	return nil
}

func (s server) AllStream(allSteam proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		// 服务端接收数据
		for {
			recv, err := allSteam.Recv()
			if err != nil {
				fmt.Println("服务端接收数据失败：", err)
				break
			}
			fmt.Printf("接收到的数据：%v\n", recv.Data)
		}
	}()
	go func() {
		defer wg.Done()
		i := 0
		// 只要调用 就向客户端发送数据  用for简单演示
		for {
			if i++; i > 10 {
				fmt.Println("服务端发送数据结束")
				break
			}
			_ = allSteam.Send(&proto.StreamRspData{
				Reply: fmt.Sprintf("%v---服务端第%d次发送", time.Now().Format("2006年01月02日 15:04:05"), i),
			})
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
	fmt.Println("服务端数据处理结束")
	return nil
}

// TODO 服务端整体流程
// 1、通过net.Listen()监听端口 获取listener
// 2、通过grpc.NewServer()创建服务 获取server
// 3、通过proto.RegisterGreeterServer()注册服务
// 4、通过server.Serve()启动服务
// 5、接收：stream.Recv()接收返回对应的结构体，发送：stream.Send()发送对应的结构体
func main() {
	// 1、监听
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	// 关闭监听
	defer listener.Close()
	// 2、创建 grpc 服务
	s := grpc.NewServer()
	// 3、注册服务
	proto.RegisterGreeterServer(s, &server{})
	// 4、启动服务
	fmt.Println("grpc server start ...")
	err = s.Serve(listener)
	if err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}

}
