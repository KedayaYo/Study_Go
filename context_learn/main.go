package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 监听cpu信息
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		// 获取cpu信息
		select {
		case <-ctx.Done():
			fmt.Println("监控退出，停止了...")
			return
		default:
			fmt.Println("监控中，CPU信息：(芯片)：", ctx.Value("chip"))
			time.Sleep(2 * time.Second)
		}

	}
}

var wg sync.WaitGroup

// 上下文
func main() {
	wg.Add(1)
	// 父上下文：context.Background() 和 context.TODO() 是最常用的两个上下文
	// context包含了主要的函数：WithCancel、WithValue、WithTimeout、WithDeadline
	// 父context结束了 子context也会结束
	//ctx1, cancel1 := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx1)
	//go cpuInfo(ctx2)
	//time.Sleep(6 * time.Second)
	//cancel1() // 一调用cancel 就会相应ctx.Done()

	// WithTimeout 主动超时	这种方式不如上面的灵活  上面的比较常用 根据逻辑去判断
	// WithDeadline 在某个时间点超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//go cpuInfo(ctx)

	valueCtx := context.WithValue(ctx, "chip", "Apple M3 Max")
	go cpuInfo(valueCtx)
	wg.Wait()
	fmt.Println("main over...")
}
