package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1、拨号
	conn, _ := net.Dial("tcp", "localhost:8020")
	// 定义一个指针
	var reply = new(string)
	// 2、使用jsonrpc创建client
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// 底层实现是不一样 返回的是一个json
	err := client.Call("HelloService.Hello", "world", reply)
	if err != nil {
		panic("调用失败")
	}
	// 3、处理调用结果
	fmt.Println(*reply)

}
