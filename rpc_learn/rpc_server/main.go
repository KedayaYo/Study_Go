package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, reply *string) error {
	// 返回值通过指针修改reply的值
	*reply = "hello:" + req
	return nil

}
func main() {
	// 1、实例化一个server
	listener, _ := net.Listen("tcp", ":8020")
	// 2、注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", new(HelloService))
	// 3、启动服务
	conn, _ := listener.Accept() // 当一个新的连接进来的时候  返回一个套接字  将套接字给rpc处理即可
	rpc.ServeConn(conn)

}
