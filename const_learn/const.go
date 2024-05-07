package main

import "fmt"

func main() {
	// 常量，定义的时候必须赋值，不能修改，常量名建议大写
	const PI float32 = 3.1415926 // 显示定义
	const PI2 = 3.1415926        // 隐式定义
	// 定义多个常量
	const (
		STATUS_OK  = 200
		STATUS_ERR = 500
		UNKNOW     = 100
		MAKE       = 1
		FEAMLE     = 2
	)

	/**
	 * 常量类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	 * 不曾使用的常量，在编译的时候会报错
	 * 显示指定类型的时候，必须确保常量的值在指定类型范围内
	 * 如果没有设置常量的值，那么默认和上一个常量的值一样
	 *
	 */

	const (
		x int = 1
		y
		s = "tth"
		z
	)
	fmt.Println(x, y, s, z)

	// iota 常量生成器
	/*
		如果中断了iota那么必须显示的恢复，后续自动递增
		自增类型默认是int
		iota可以建行const定义
		每次出现const的时候iota就会重置为0
	*/
	const (
		a = iota // 0
		b = iota // 1
		c        // 2
		d = "hh"
		e = "hh"
		f = 100
		g = iota
	)
	fmt.Println(a, b, c, d, e, f, g)
}
