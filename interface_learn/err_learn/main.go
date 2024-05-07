package main

import "fmt"

func mPoint(datas ...interface{}) {
	for i, v := range datas {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}
}

func nPoint(data interface{}) {
	fmt.Printf("value: %v\n", data)
}

// 实现error接口
type MyError struct {
}

func (e *MyError) Error() string {
	fmt.Println("Error() 方法被调用")
	return "这是自定义Error"
}

// 常见的错误
func main() {
	// 空结构体 装万物 但是注意空结构体作为函数参数传递的时候的坑
	//var data = []interface{}{
	//	"Kedaya", 20, 180,
	//}
	//mPoint(data...)

	// 这个时候为什么data...报错呢：想要强行将切片类型打散给 []interface{}是不行的
	var data = []string{
		"Go", "Python", "Java",
	}

	// ... 打散slice
	nPoint(data)

	// 如果想要将切片类型打散给 []interface{} 可以使用以下方式  转换一次
	var datai = []interface{}{}
	for _, v := range data {
		datai = append(datai, v)
	}
	mPoint(datai...)

	var err error
	err = &MyError{}
	res := err.Error()
	fmt.Println(res)
}
