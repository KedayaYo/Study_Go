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

// 结构体嵌套
func main() {
	stu1 := Student{
		Person{
			"Entic", 18,
		},
		"Kedaya",
		99.8,
	}
	//fmt.Printf("stu1: %v\n", stu1)
	// 第一种嵌套取值按层级比较麻烦
	//fmt.Printf("stu1 name: %v\n", stu1.p.name)
	// 第二种嵌套取值比较简单 但是会取最外层的值  赋值也是一样
	stu1.name = "小火龙"
	fmt.Printf("stu1: %v\n", stu1)
	fmt.Printf("stu1 name: %v\n", stu1.name)
}
