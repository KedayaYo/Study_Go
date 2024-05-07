package main

import (
	"context"
	"fmt"
	proto "go_learn/grpc_token_auth_learn/proto/pg_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": "homeland",
		"uid":   "1001",
	}, nil
}

// 传输是否安全 暂时给false即可
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 1、创建intercept
	// func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {}
	helloUnaryIntercept := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("客户端拦截到了一个请求%s\n", method)
		start := time.Now()

		// 方式1:使用metadata传递token等信息
		//mdSlice := make([]string, 0, 2)
		//mdSlice = append(mdSlice, "token")
		//mdSlice = append(mdSlice, "homeland")
		//mdSlice = append(mdSlice, "uid")
		//mdSlice = append(mdSlice, "11001")
		//ctx = metadata.AppendToOutgoingContext(ctx, mdSlice...)
		// 直接在函数调用中构造并传递键值对
		//ctx = metadata.AppendToOutgoingContext(ctx, "token", "homeland", "uid", "11001")

		// 方式2:使用grpc中内置的WithPerRPCCredentials(返回是一个接口type PerRPCCredentials interface 实现里面的两个方法即可)
		// 本质也是metadata 服务端还是一样的
		grpc.WithPerRPCCredentials(customCredential{})

		// 执行初始的逻辑 invoker(ctx, method, req, reply, cc, opts...)
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("客户端耗时：%v\n", time.Since(start))
		return err
	}
	// 配置集合
	opts := make([]grpc.DialOption, 0, 2)
	//opts := []grpc.DialOption{}
	// 2、grpc生成客户端连接配置：WithUnaryInterceptor
	opts = append(opts, grpc.WithUnaryInterceptor(helloUnaryIntercept))
	// 3、grpc 连接
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Printf("%v\n", opts)
	conn, err := grpc.Dial("localhost:8020", opts...)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer conn.Close()
	// 4、实例化客户端
	client := proto.NewHelloServiceClient(conn)
	response, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name:  "Kedaya",
		Age:   18,
		Hobby: []string{"Game", "Drink"},
	})
	if err != nil {
		panic("failed to call: " + err.Error())
	}
	fmt.Println(response.Reply)
}
