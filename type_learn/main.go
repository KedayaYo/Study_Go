package main

import (
	"fmt"
	"strconv"
)

type MyInt1 int // 类型定义 需要强制转换  这是自定义类型和int不一样  一般用于基本类型的封装 比如在int类型上加一个string 转换的方法
// 这是一个方法
func (mi MyInt1) string() string {
	return strconv.Itoa(int(mi))
}

type MyInt2 = int // 类型别名
func main() {
	// type关键字
	/*
		1、定义结构体
		2、定义接口
		3、定义别名
		4、类型定义
		5、类型判断
	*/
	// 别名实际上是为了更好的理解代码
	var myint1 MyInt1 = 1
	var myint2 MyInt2 = 1
	a := 10
	fmt.Printf("%T\n", myint1) // main.MyInt
	fmt.Printf("%T\n", myint2) // int
	fmt.Println(a + int(myint1))
	fmt.Println(a + myint2) // 11 在编译的时候 会直接替换成int类型
	// 在类型定义中封装方法
	fmt.Println(myint1.string())

	// 类型判断
	// 方法1: switch .
	var x interface{} = "hello"
	switch x.(type) {
	case string:
		fmt.Println("string")
	}
	//y := x.(string) // 类型断言 转换成string 后续会展示
	//fmt.Println(y)

}
