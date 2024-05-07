package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	fmt.Println("async print")
}
func main() {
	fmt.Printf("main goroutine start\n")
	// 主死随从
	//go asyncPrint()
	// 匿名函数启动goroutine
	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("goroutine print")
	//	}
	//}()

	// 1、闭包 2、for循环的问题 for循环每个变量会重用
	// 每次for循环的时候 i变量会被重用  当我进入第二轮for循环的时候 这个i就变了
	//go func() {
	//			fmt.Println(i)
	//		}()
	for i := 0; i < 100; i++ {
		// 第一种解决办法：temp:=i  不保证顺序
		//temp := i
		//go func() {
		//	fmt.Println(temp)
		//}()

		// 第二种解决办法：传参  值传递  自己复制i
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(3 * time.Second)
}
