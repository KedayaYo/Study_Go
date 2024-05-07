package main

import (
	"fmt"
)

// 函数定义 func 函数名 参数列表(参数 参数类型) 返回值列表(返回值 返回值类型) {}
func add(a int, b int) (sum int, err error) {
	fmt.Println("add function")
	//err := errors.New("add error")
	err = fmt.Errorf("add error")
	sum = a + b
	return sum, err
}

func addition(item ...int) (sum int) {
	for _, v := range item {
		sum += v
	}
	return sum
}

// 省略号 ... 表示不定参数  可以不传  很有用
func judge(item ...int) {
	fmt.Println("judge function")
	// item ...int相当于一个切片 []int
	for _, v := range item {
		fmt.Println(v)
	}
}

func cel(op string, o int, a int, b int, myfunc func(item ...int) int) (res int) {
	switch op {
	case "+":
		fmt.Println("加法")
		return myfunc(a, b) + o
	case "-":
		fmt.Println("减法")
		return myfunc(a, b) - o
	case "*":
		fmt.Println("乘法")
		return myfunc(a, b) * o
	case "/":
		fmt.Println("除法")
		return myfunc(a, b) / o
	default:
		fmt.Println("op 参数错误")
	}
	return 0
}

func logFunc(s string) {
	fmt.Println("日志函数: " + s)
}

func callback(s string, myfunc func(s string)) {
	myfunc("bbb")
	fmt.Println("回调函数: " + s)
}

var local int

func autoIncrement() int {
	local += 1
	return local
}

func autoIncrementClosure() func() int {
	closure := 0
	// 如果要返回int  直接写函数代表返回函数  和返回类型不对应  需要加一个() 表示调用
	//return func() int {
	//	closure += 1
	//	return closure
	//}()

	// 但是一个函数访问另一个函数的局部变量是不行的  但是在一个函数中定义使用匿名函数就可以 这就是闭包

	// 使用匿名函数调用  不满足当前需求  应该把函数返回出去
	return func() int {
		closure += 1
		return closure
	}
}

func main() {
	// Go 函数支持普通函数、匿名函数、闭包
	/*
		Go中函数是“一等公民”
			1、函数本身可以作为变量
			2、匿名函数 闭包
			3、函数可以满足接口
	*/
	a := 10
	b := 20
	c := 30
	sum, _ := add(a, b)
	fmt.Println(sum)
	judge(a, b, c)

	// 将函数作为变量  只需要传递函数名即可  传递调用  声明的就是返回值类型而不是函数类型
	funcVar := add
	fmt.Printf("%T\n", funcVar)

	res := cel("+", 10, 1, 2, addition)
	fmt.Println(res)
	// 一般用作回调
	callback("aaa", logFunc)

	// 匿名函数 func()
	localFunc := func(s string) {
		fmt.Println("匿名函数: " + s)
	}
	callback("ccc", localFunc)

	// 闭包
	// 闭包是一个函数，这个函数包含了它外部作用域的一个变量
	// 有一个需求：每次调用的时候  返回结果都加一

	// 方法1: 定一个全局变量  但是不推荐  这种方法被迫定义了一个全局变量
	//for i := 0; i < 5; i++ {
	//	fmt.Println(autoIncrement())
	//}

	// 方法2：使用闭包
	closureFunc := autoIncrementClosure()
	for i := 0; i < 5; i++ {
		// 每次调用只会执行匿名函数里面的内容  并不会调用autoIncrementClosure的变量了
		fmt.Println(closureFunc())
	}
}
