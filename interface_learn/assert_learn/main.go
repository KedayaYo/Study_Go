package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addInterfaceInt(a, b interface{}) int {
	// 断言转换格式：value, ok := interface.(type)
	// value: 断言成功后的值
	// ok: 断言是否成功
	// interface: 需要断言的接口
	// type: 断言的类型
	// 因为不知道会传入什么值进来  转换还可能失败  需要判断
	ai, ok := a.(int)
	if !ok {
		panic("a is not int")
	}
	bi, ok := b.(int)
	if !ok {
		panic("b is not int")
	}
	return ai + bi
}

// 通过switch 来断言
func addInterface(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case int64:
		ai, _ := a.(int64)
		bi, _ := b.(int64)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("not supported type")
	}
}

// 断言
func main() {
	a := 1
	b := 2
	res := addInterface(a, b)
	fmt.Println(res)
}
