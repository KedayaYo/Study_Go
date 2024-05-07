package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, reply *string) error {
	// 返回值通过指针修改reply的值
	*reply = "hello:" + req
	return nil

}
func main() {
	// 1、注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", new(HelloService))
	// 2、实例化一个server
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.Reader
			io.Closer
		}{
			Writer: w,
			Reader: r.Body,
			Closer: r.Body,
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})
	// 3、监听端口启动
	_ = http.ListenAndServe(":8020", nil)

}
