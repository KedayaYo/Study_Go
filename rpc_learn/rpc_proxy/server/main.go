package main

import (
	"go_learn/rpc_learn/rpc_proxy/handler"
	"go_learn/rpc_learn/rpc_proxy/proxy"
	"net"
	"net/rpc"
)

func main() {
	// 1、实例化一个server
	listener, _ := net.Listen("tcp", ":8020")
	// 2、注册处理逻辑 handler
	proxy.RegisterHelloService(handler.HelloServiceName, &handler.HelloService{})
	for {
		// 3、启动服务
		conn, _ := listener.Accept() // 当一个新的连接进来的时候  返回一个套接字  将套接字给rpc处理即可
		go rpc.ServeConn(conn)
	}

}
