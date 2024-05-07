/**
 * @Author Kedaya
 * @Date 2024/5/6 00:56:00
 * @Desc
 **/
package handler

const HelloServiceName = "handler/HelloService"

type HelloService struct {
}

func (s *HelloService) Hello(req string, reply *string) error {
	// 返回值通过指针修改reply的值
	*reply = "hello:" + req
	return nil

}
