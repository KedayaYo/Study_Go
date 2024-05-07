/**
 * @Author Kedaya
 * @Date 2024/5/6 00:58:00
 * @Desc
 **/
package proxy

import (
	"go_learn/rpc_learn/rpc_proxy/handler"
	"net/rpc"
)

type HelloService struct {
	*rpc.Client
}

func (h *HelloService) Hello(req string, reply *string) error {
	return h.Client.Call(handler.HelloServiceName+".Hello", req, reply)
}

// 需要初始化一个Client
func NewHelloServiceClient(protcol, address string) (HelloService, error) {
	client, err := rpc.Dial(protcol, address)
	if err != nil {
		return HelloService{}, err
	}
	return HelloService{Client: client}, nil
}
