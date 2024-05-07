package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"time"
)

type ResponseData struct {
	Data int `json:"data"`
}

func AddRequest(a, b int) int {
	req := HttpRequest.NewRequest()
	req.SetTimeout(60 * time.Second)
	res, _ := req.Get(fmt.Sprintf("http://localhost:8000/add?a=%d&b=%d", a, b))
	body, _ := res.Body()
	//fmt.Printf("result: %s\n", string(body))
	responseData := ResponseData{}
	json.Unmarshal(body, &responseData)
	return responseData.Data
}

func main() {
	res := AddRequest(2, 2)
	fmt.Printf("result: %d\n", res)
}
