package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pg "go_learn/rpc_learn/proto/pg_go"
)

type HelloWorld struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	req := &pg.HelloRequest{
		Name:  "Kedaya",
		Age:   22,
		Hobby: []string{"basketball", "football"},
	}
	// 序列化 proto.Marshal入参是一个接口  返回值是[]byte切片
	rsp, _ := proto.Marshal(req) // 具体编码是如何做到的  protobuf的原理 varint变长的int
	fmt.Printf("%v\n", len(rsp)) // Kedaya 32
	// 反解 1、定义一个空的请求 2、proto.Unmarshal(rsp, decodeReq)
	decodeReq := &pg.HelloRequest{}
	proto.Unmarshal(rsp, decodeReq)
	fmt.Printf("%v\n", decodeReq) // name:"Kedaya"  age:22  hobby:"basketball"  hobby:"football"

	// 对比json
	jsonStruct := HelloWorld{
		"Entic",
		20,
		[]string{"music", "movie"},
	}
	jsonRsp, _ := json.Marshal(jsonStruct)
	fmt.Printf("%v\n", len(jsonRsp)) // {"name":"Entic"}  52  基本类型的json编码是protobuf的两倍
	// 反解
	decodeJsonReq := &HelloWorld{}
	json.Unmarshal(jsonRsp, decodeJsonReq)
	fmt.Printf("%v\n", *decodeJsonReq) // {Entic 20 [music movie]}
}
