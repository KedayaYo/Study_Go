package main

import (
	"fmt"
	"sync"
)

// 子 goroutine 退出时，会向父 goroutine 发送信号，父 goroutine 会收到这个信号，然后继续执行。
func main() {
	// 定义WaitGroup 不需要实例化 内部是一个struct
	var wg sync.WaitGroup
	// Add 监控多少个携程
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			// Done 通知WaitGroup，当前携程执行完毕 Add和Done一起出现
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	// Wait 会阻塞，直到所有的携程执行完毕
	wg.Wait()
	fmt.Println("over")

	// waitgroup 主要用于携程的执行等待 Add和Done配合使用
}
