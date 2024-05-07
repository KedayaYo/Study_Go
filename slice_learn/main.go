package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	// slice初始化
	//var s []int  // 没有初始化的slice是nil  无法直接赋值
	//var s := []int{1, 2, 3, 4, 5}
	//s := make([]int, 5)
	// 左闭右开
	//s = s[0:2]
	// 从最开始到下标2
	//s = s[:3]
	// 从下标1到最后
	//s = s[1:]

	// 未初始化的slice是nil  不能直接赋值  但是可以通过append扩容
	var arr1 []string
	arr1 = append(arr1, "a", "b", "c")
	//arr1[0] = "a" // panic: runtime error: index out of range [0] with length 0
	//arr1[0] = "a"
	//arr1[1] = "b"
	//arr1[2] = "c"
	fmt.Println(arr1)

	// 两个slice拼接
	// 方法1: for循环
	sl1 := []string{"Go", "Python", "Java"}
	sl2 := []string{"C", "C++", "Ruby"}
	for _, v := range sl2 {
		sl1 = append(sl1, v)
	}
	fmt.Println(sl1)
	// 方法2: ...运算符打散
	//sl1 = append(sl1, sl2...)
	// 也可以直接切片分割
	sl1 = append(sl1, sl2[1:]...)
	fmt.Println(sl1)

	// 删除 比较麻烦 需要取前面的元素和后面的元素进行拼接
	myslice := append(sl2[:1], sl2[2:]...)
	fmt.Println(myslice)

	// 复制 生成一个新的slice  将右边的copy到左边
	// var copySlice []string  这样声明并不会有空间开辟  copy也不会扩容  返回的是一个空的slice
	copySlice := make([]string, len(myslice))
	copy(copySlice, myslice)
	fmt.Println(copySlice)

}
