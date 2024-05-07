package main

import "fmt"

func deferReturn() (res int) {
	defer func() {
		res++
	}()
	return 10
}

func main() {
	// 连接数据库、打开文件、开始锁 .... 无论如何都要关闭资源 相当于其他语言finally
	// 用锁举例
	//var mu sync.Mutex
	//mu.Lock()
	// defer后面的代码会放在函数return之前执行
	//defer mu.Unlock()

	// 多个defer的执行顺序是后进先出  会压栈  main -> 2 -> 1
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//fmt.Println("main")

	// defer 是有能力去修改返回值的
	// return可以分为两步：1、给返回值赋值 2、RET指令  defer是在给返回值赋值之后，RET指令之前执行
	res := deferReturn()
	fmt.Println(res) // 11
}
