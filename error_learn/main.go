package main

import (
	"fmt"
)

func A() (int, error) {
	// 初始化error
	//err := errors.New("error New")
	err := fmt.Errorf("fmt error")
	return 0, err
}

func recoverFunc() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func main() {
	// Go 语言错误处理理念：一个函数可能出错 但是不要返回错误，而是返回一个错误值
	// Go 语言错误处理方式：返回值 + err 因此代码中会出现大量的if err != nil

	if _, err := A(); err != nil {
		fmt.Println(err)
	}

	// panic 是一个函数  会导致程序崩溃退出  不推荐使用
	// 一般用于程序启动期间  依赖于其他的服务、mysql连接等
	// panic 会导致程序崩溃退出，但是panic之前的defer语句会在panic之前执行
	//defer fmt.Printf("defer1....")
	//panic("this is a panic")
	//defer fmt.Printf("defer2....")
	//fmt.Println("....")

	// 程序触发panic panic: assignment to entry in nil map  没有初始化map
	// recover 可以捕获panic
	// 1、recover只有在defer的函数中使用才会生效 2、defer必须在panic之前定义
	// 3、recover只能捕获最后一个panic 4、recover处理异常后逻辑并不会恢复到panic那个点
	// 5、recover返回的是interface{}类型 6、多个defer会按照 先进后出的顺序执行  形成栈
	// recoverFunc()不可以提到其他函数去掉用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover: ", err)
		}
	}()

	var maps map[string]string
	maps["name"] = "Kedaya"
}
