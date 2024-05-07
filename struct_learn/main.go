package main

import "fmt"

type Person struct {
	Name string
	Age  int
	gen  string
}

func main() {
	// 需求：需要放多个信息到person中
	// 1、使用切片
	//var personArr [][]string
	//personArr:=append(personArr, []string{"Entic", "20", "男"})
	// 2、使用interface
	//var personArr [][]interface{}
	//personArr:=append(personArr, []interface{}{"Entic", 20, "男"})// 断言  可以直接放入不同类型 之后通过断言转换
	// 2、使用结构体
	// 如何初始化结构体
	// 方法1：直接初始化
	p1 := Person{"Entic", 20, "男"}
	fmt.Printf("%#v\n", p1)
	// 更灵活 不需要全部赋值也可以初始化
	// 方法2:使用键值对
	p2 := Person{
		Name: "Bing",
		Age:  22,
		//gen:  "女",
	}
	fmt.Printf("%#v\n", p2)
	// 方法3：使用切片
	//var persons []Person
	//persons = append(persons, p1, p2)
	//persons:= []Person{p1, p2}
	persons := []Person{
		{"Entic", 20, "男"},
		{
			Name: "Bing",
		},
	}
	fmt.Printf("%#v\n", persons)
	// 方法4：赋值
	var p Person
	p.Name = "Entic"
	fmt.Printf("%#v\n", p)

	// 匿名结构体：用于一些临时场景  是一次性的 比如将地址信息整合到一个结构体中
	address := struct {
		Contry   string
		Province string
		City     string
	}{
		Contry:   "中国",
		Province: "上海",
		City:     "上海",
	}
	fmt.Printf("%#v\n", address.Province)
}
