package main

import "fmt"

func main() {
	// slice原理
	// go的slice在函数参数传递的时候是值传递  效果呈现出引用效果
	// 但是通过append函数修改slice的时候会重新分配内存  不会再跟原来的slice有关联
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:6]
	// 长度是5 但是容量是9  后面分配的容量还在
	fmt.Println(cap(s1))
	fmt.Println(len(s1))
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(cap(s1)) // 18  即使扩容变成新的slice 后面分配的容量还在
	fmt.Println(len(s1))
	s2 := data[2:7]
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}
