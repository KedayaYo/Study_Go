package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1、连接远程rpc服务
	client, _ := rpc.Dial("tcp", "localhost:8020")
	// 2、调用远程方法(HelloService.Hello 需要补全作用于HelloService来.出方法)
	// 定义一个指针
	var reply = new(string)
	err := client.Call("HelloService.Hello", "world", reply)
	if err != nil {
		panic("调用失败")
	}
	// 3、处理调用结果
	fmt.Println(*reply)

}
