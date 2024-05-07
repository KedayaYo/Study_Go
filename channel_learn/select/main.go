package main

import (
	"fmt"
	"time"
)

func g1(ch1 chan struct{}) {
	time.Sleep(time.Second)
	// 声明struct{} 实例化struct{}{} 一个空的chan  来传递信号 过了1s 告诉主程序的chan接收到了一个空的struct 就结束了主程序
	ch1 <- struct{}{}
}
func g2(ch2 chan struct{}) {
	time.Sleep(2 * time.Second)
	ch2 <- struct{}{}
}

var done = make(chan struct{})

func main() {
	// select 语句类似于switch语句 但是select的功能和我们操作linux里面提供io的select、poll、epoll差不多
	// select主要作用于多个channel
	// 需求：有两个goroutine都在执行  但是主goroutine中  当某个执行完成后我马上知道
	// channel是线程安全的   多个goroutine往一个channel写数据是安全的
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go g1(ch1)
	go g2(ch2)
	//<-done
	// 通过select监控多个channel
	// 1、某一个分支就绪了就执行该分支 2、如果两个都就绪了  先执行哪个随机  目的是：防止饥饿（防止一直执行某一个）
	// 应用场景：我不知道channel会运行多久 如果某个channel卡住了  就会导致主程序结束不了
	// 方法1:可以通过for+select的default进行一个超时机制 default会直接不等待这个时候加上我们自己的定的超时时间即可
	// 方法2:通过time.NewTimer .C 返回的也是一个channel 根据这个channel来判断即可
	// select 是阻塞的  放在go携程的方法里面的时候 会阻塞select后面的代码
	tc := time.NewTimer(500 * time.Millisecond)
	for {
		select {
		case <-ch1:
			fmt.Println("g1执行完成")
		case <-ch2:
			fmt.Println("g2执行完成")
		case <-tc.C:
			fmt.Println("超时")
			return
			//default:
			//	time.Sleep(500 * time.Millisecond)
			//	fmt.Println("超时")
			//return
		}
	}
}
