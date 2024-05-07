package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan int, 2)
	go func(msg chan int) {
		for data := range msg {
			fmt.Println("接收到的数据", data)
		}
		fmt.Println("通道关闭")
	}(msg)

	msg <- 1
	msg <- 2
	close(msg) // 关闭掉channel 关闭之后  上面的for循环就会结束  已经关闭的channel不能再写入数据但是可以继续取值
	//msg <- 3
	//data:=<-msg
	time.Sleep(5 * time.Second)
}
