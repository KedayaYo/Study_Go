package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 通信
func main() {
	// 不要通过共享内存来进行通信，而应该通过通信来共享内存
	// 定义
	// var 变量 chan 元素类型
	// make(chan 元素类型, [缓冲大小])
	// var 变量 chan<- 元素类型 // 只能发送
	// var 变量 <-chan 元素类型 // 只能接收
	// 通道是引用类型，必须使用make函数初始化
	//var msg chan string
	//var msg = make(chan string, 0) // 适用于通知  B第一时间知道A是都已经完成
	//var msg = make(chan string, 1)// 适用于消费者和生产者之间的通信
	// 接收
	//msg <- "hello"
	// 读取
	//data := <-msg
	//fmt.Println(data)
	var msg = make(chan string, 0) // channel的初始化值 如果为0  放值进去会阻塞
	// 无缓冲的通道 读写都会阻塞  通过goroutine 来解决
	wg.Add(1)
	go func(msg chan string) { // Go有一种 happen-before 机制 可以保障
		defer wg.Done()
		data := <-msg
		fmt.Println("接收到的数据", data)
	}(msg)
	msg <- "hello"
	wg.Wait()
	// waitGroup 如果少了done调用 容易出现deadlock，无缓冲的channel也会出现

	/*
		go 中channel的应用场景
		1、消息传递、消息过滤
		2、信号广播
		3、事件订阅和广播
		4、任务分发
		5、结果汇总
		6、并发控制
		7、同步和异步
		...
	*/
}
