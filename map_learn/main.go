package main

import (
	"fmt"
)

func main() {
	// map是一个key（所索引）和 value（值）的无序集合  方便查询
	// 定义 var 变量名 map[key类型]value类型
	// 当前定义是一个nil map  赋值的时候必须初始化var courseMap map[string]string{}  可以使用make
	// new 和 make的区别：new返回的是指针，make返回的是初始化后的对象
	//var courseMap map[string]string
	//var courseMap map[string]string{}
	//var courseMap = make(map[string]string,3)

	// map 不是线程安全的  在携程中使用需要使用：
	//var syncMap sync.Map
	//syncMap := sync.Map{}
	var courseMap = map[string]string{
		"Go":     "Go语言",
		"Java":   "Java语言",
		"Python": "Python语言",
	}
	// 取值
	fmt.Println(courseMap["Go"])

	// 赋值
	courseMap["C++"] = "C++语言"
	fmt.Printf("%v\n", courseMap)

	// 删除 不存在元素不会报错
	delete(courseMap, "Java")
	fmt.Printf("%v\n", courseMap)

	// 判断是否存在
	if _, ok := courseMap["Java"]; ok {
		fmt.Println("Java存在")
	} else {
		fmt.Println("Java不存在")
	}

	// 遍历
	for key, value := range courseMap {
		fmt.Println(key, value)
	}
}
