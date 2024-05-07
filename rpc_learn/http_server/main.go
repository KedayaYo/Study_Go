package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http://localhost:8080/add?a=1&b=2
	// 返回json格式：{"data": 3}
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// 解析参数
		_ = r.ParseForm()
		// 打印请求路径
		fmt.Println("path: ", r.URL.Path)
		// 取get参数
		a, _ := strconv.Atoi(r.Form.Get("a"))
		b, _ := strconv.Atoi(r.Form.Get("b"))
		// 设置为json格式
		w.Header().Set("Content-Type", "application/json")
		// 返回方法
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jData)
	})

	// 监听端口 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Failed to start http_server: ", err)
	}
}
