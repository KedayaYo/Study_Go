package main

import (
	"fmt"
	"go_learn/rpc_learn/rpc_proxy/proxy"
)

func main() {
	// 1、连接远程rpc服务
	client, _ := proxy.NewHelloServiceClient("tcp", "localhost:8020")
	// 2、调用远程方法(HelloService.Hello 需要补全作用于HelloService来.出方法)
	// 定义一个指针
	var reply = new(string)
	client.Hello("world", reply)
	// 3、处理调用结果
	fmt.Println(*reply)

}
