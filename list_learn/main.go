package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 空间不连续 空间浪费(因为里面会生成一个指针指向下一个元素，这样才知道下一个元素是什么)
	// 性能差异比较大  slice查询比较快 list插入和删除比较快 用的比较少
	// 优点：1、可以提前知道容量 2、可以提前知道长度
	// 定义
	//var myList list.List
	//myList := list.List{}
	//myList := list.New()
	myList := list.New()
	myList.PushBack("Go")
	myList.PushBack("Python")
	myList.PushBack("Java")
	fmt.Println(myList)

	// 在头部放数据
	myList.PushFront("C++")

	// 在指定元素之后插入
	myList.InsertAfter("C", myList.Front())

	// 在指定元素之前插入
	myList.InsertBefore("C#", myList.Back())

	// 在中间插入
	e := myList.Front()
	for ; e != nil; e = e.Next() {
		// 在遍历时插入
		if e.Value.(string) == "Java" {
			break
		}
	}
	myList.InsertBefore("PHP", e)

	// 删除元素
	myList.Remove(e)

	// 遍历打印
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 反向遍历打印
	for e := myList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
