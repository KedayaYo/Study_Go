/**
 * @Author Kedaya
 * @Date 2024/5/6 01:08:00
 * @Desc
 **/
package proxy

import (
	"net/rpc"
)

// 相当于让HelloService实现了Hi接口
type HiService interface {
	Hello(req string, reply *string) error
}

func RegisterHelloService(serviceName string, srv HiService) error {
	return rpc.RegisterName(serviceName, srv)
}
