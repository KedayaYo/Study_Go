package main

import "fmt"

type Person struct {
	name string
	age  int
}

func modifyName(p *Person, name string) {
	p.name = name
}

// 通过指针交换两个值
func swap(a, b *int) {
	// 这样是错误的
	//a, b = b, a
	// 原因：给a,b分配空间的时候  会分配地址  通过指针获取的时候 是通过地址获取的值
	// 而方法这边是值传递  也就是复制了一份地址过来  所以在方法内部改变了地址的位置  但是并没有改变a,b的值
	// 正确做法是 将a的地址的值改为b地址的值  将b的地址的值改为a地址的值
	t := *a
	*a = *b
	*b = t
}

func main() {
	p := Person{"Entic", 20}
	modifyName(&p, "Kedaya")
	fmt.Printf("p:%#v\n", p)

	// 指针的定义
	//var poi *int
	//var po *Person
	//po = &p
	// 初始化
	po := &Person{
		name: "Entic",
		age:  20,
	}
	fmt.Printf("po:%p\n", po)
	fmt.Printf("po:%#v\n", *po)

	// 取值赋值
	// Go：可以使用指针类型的变量来访问其所指向的变量再来操作  也可以直接用指针变量来操作
	// Go: go的指针限制了指针的运算 但是可以使用unsafe包来操作指针
	(*po).name = "Kedaya"
	po.age = 21
	fmt.Printf("po:%#v , po.name:%#v\n", *po, po.name)

	// 初始化 初始值是nil 在变量上取地址  而不是var a int；b:=&5
	//var a int = 10
	//b := &a
	// 初始值是nil panic: runtime error: invalid memory address or nil pointer dereference
	// 需要初始化
	//var pp *Person
	//fmt.Printf("p:%#v\n", pp.name)
	// 初始化方式1:
	//pp := &Person{}
	// 初始化方式2:
	//pp := new(Person)
	// 初始化方式3: 对于结构体会自己初始化
	var originPerson Person
	pp := &originPerson
	fmt.Printf("pp:%#v\n", pp.name)

	// 初始化的两个关键字：make和new
	// make：用于内建类型（map、slice 和channel）的内存分配 返回的是这三个引用类型本身
	// new：用于各种类型的内存分配，返回的是指针
	// map和指针必须初始化

	// 通过指针交换两个值
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a:%d, b:%d\n", a, b)
}
