package main

import (
	"fmt"
	"log"
)

// 全局变量和局部变量
// var name = "Entic"
// var age = 20
var (
	name = "Entic"
	age  = 20
)

type MyError struct {
	code int64
	msg  string
}

// 实现error接口的Error方法
func (e MyError) Error() string {
	return fmt.Sprintf("Error: [%d] %s", e.code, e.msg)
}

func (e MyError) Code() int64 {
	return e.code
}

func (e MyError) Msg() string {
	return e.msg
}

func a() (int, bool) {
	return 0, false
}

func main() {
	// Go 是静态语言，变量的类型是在编译期确定的
	// 1、变量必须先定义 2、变量必须又类型 3、类型定下来不能改变

	// 声明变量
	//var name string
	//name = "Kedaya"
	name := "Kedaya"
	fmt.Println(name)

	// 声明多个变量
	//var user1, user2, user3 string
	var user1, user2, user3 = "Entic", true, 110
	fmt.Println(user1, user2, user3)
	/**
	注意：
		变量必须先定义，再使用
		go语言是静态语言，要求变量的类型和值必须一致
		变量名不能重复定义
		简介变量定义方式，只能在函数内部使用
		变量是有零值的，int 0，string ""
	*/

	// 匿名变量
	// 匿名变量用一个下划线_表示，匿名变量不占用内存空间，不会分配内存
	var _ int
	// 由于我当前不需要res的返回值  只需要ok判断是否成功
	// 这个时候就可以用同类型的匿名变量  不然就需要打印res 不让他报错
	_, ok := a()
	if true != ok {
		//fmt.Println(res)
		err := MyError{
			code: 100,
			msg:  "调用错误",
		}
		//fmt.Printf("报错: %v", err)
		log.Printf("报错: %v", err)
	}
}
