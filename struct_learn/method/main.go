package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	// 第一种嵌套方式
	//p     Person
	// 第二种嵌套方式：匿名嵌套
	Person
	name  string
	score float32
}

func stuPrintFunc(p Person) {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

func (p *Person) stuPrintMethod1() {
	p.age = 19
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

func (p Person) stuPrintMethod2() {
	p.age = 20
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

// 结构体定义方法
func main() {
	// 方法和函数的区别
	// 函数：独立的功能，独立的代码块
	// 方法：和结构体绑定的函数，方法是一个特殊的函数
	// 方法的定义：func (接收者变量 接收者类型) 方法名(参数列表) (返回值列表) {代码块}
	// 接收者变量：类似于函数的形参，表示调用该方法的具体实例
	// 接收者类型：接收者的类型，可以是指针类型和非指针类型
	// 方法只能和对应的类型绑定，也就是说，方法只能在定义的类型上调用
	// 接收者类型最好定义为指针类型，这样可以避免进行值传递

	stu1 := Student{
		Person{
			"Entic", 18,
		},
		"Kedaya",
		99.8,
	}
	stu2 := &Student{
		Person{
			"Entic", 18,
		},
		"Kedaya",
		99.8,
	}

	// Go语言的优化：针对结构体struct
	// 如果方法的接收者是指针类型，那么Go语言会自动转换为指针类型调用  不需要再将变量转换为指针类型
	// 相反同理 自动相互转换
	stuPrintFunc(stu1.Person)
	stu1.stuPrintMethod1()
	stu1.stuPrintMethod2()
	stu2.stuPrintMethod2()
	stu2.stuPrintMethod1()
}
