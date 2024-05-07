package main

import "fmt"

func main() {
	// Go 语言提供了哪些集合类型的数据结构：数组、切片、map、list
	// 数组：长度固定，类型一致 定义：var 变量名 [初始化数量]int
	// 切片：长度不固定，类型一致 定义：var 变量名 []int
	// map：键值对 定义：var 变量名 map[string]int
	// list：双向链表 定义：var 变量名 list.List
	// 数组和切片的区别：数组长度固定，切片长度不固定
	// map和list的区别：map是键值对，list是双向链表

	// 数组的类型前面是带数字的  也就是说不同大小的数组是不能相互赋值的
	//var arr1 [5]int
	//var arr2 [4]int
	//arr1 = arr2 // cannot use arr2 (type [4]int) as type [5]int in assignment
	//fmt.Printf("arr1: %T\n", arr1) //arr1: [5]int
	//fmt.Printf("arr2: %T\n", arr2) //arr2: [4]int

	var arr1 [3]string
	arr1[0] = "a"
	arr1[1] = "b"
	arr1[2] = "c"
	fmt.Println(arr1) //[a b c]

	// range 遍历数组
	for i, v := range arr1 {
		fmt.Printf("arr1[%d]: %s\n", i, v)
	}

	// 数组的初始化
	//var arr2 [3]string = [3]string{"Go", "Java", "Python"}
	//var arr2 = [3]string{"Go", "Java", "Python"}
	//arr2 := [3]string{"Go", "Java", "Python"}

	// 数据初始化的时候指定某个元素的值
	arr2 := [3]string{2: "Java"}
	fmt.Println(arr2) //[ Java]

	// 不指定长度的数组
	arr3 := [...]string{"Go", "Java", "Python"}
	fmt.Println(arr3) //[Go Java Python]

	// 数组的判断  首先长度要一样  然后每个元素的值也要一样  顺序不能乱
	arr4 := [3]string{"Java", "Go", "Python"}
	if arr3 == arr4 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	// 多维数组
	var arrInfo [3][4]string
	arrInfo[0] = [4]string{"Go", "Go语言", "gin", "mysql"}
	arrInfo[1] = [4]string{"Java", "Java语言", "spring", "mysql"}
	arrInfo[2] = [4]string{"Python", "Python语言", "flask", "mysql"}
	// 遍历
	//for i, row := range arrInfo {
	//	for j, col := range row {
	//		fmt.Printf("arrInfo[%d][%d]: %s\n", i, j, col)
	//	}
	//}

	for i := 0; i < len(arrInfo); i++ {
		for j := 0; j < len(arrInfo[i]); j++ {
			fmt.Printf("arrInfo[%d][%d]: %s\n", i, j, arrInfo[i][j])
		}
	}
}
