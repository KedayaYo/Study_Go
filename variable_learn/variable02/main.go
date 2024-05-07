package main

import (
	"log"
	"strconv"
)

func main() {
	//// 整型
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	//// 根根系统位数有关 32位系统是int32 64位系统是int64
	//var e int // 动态类型
	//
	//// 浮点型
	//var f1 float32
	//var f2 float64

	// 字符
	//var g byte // 主要适用于存放字符  是uint8 相当于别名
	//var h rune // 主要适用于存放中文字符 是int32 相当于别名
	//var i string
	//g = 'a'
	//log.Printf("%c\n", g)
	//h = '好'
	//log.Printf("%c\n", h)
	//i = "hello"
	//log.Printf("%s\n", i)

	// Go允许在底层结构相同的两个类型之间转换
	type MyInt int
	var a MyInt = 1
	// 将a(MyInt)转换为int b现在是int类型
	b := int(a)
	// 将b(int)转换为MyInt c现在是MyInt类型
	c := MyInt(b)
	log.Printf("%T %v\n", a, a)
	log.Printf("%T %v\n", b, b)
	log.Printf("%T %v\n", c, c)

	// 字符串转数字
	log.Println("字符串转数字")
	istr := "123"
	//istr := "123h"
	atoi, err := strconv.Atoi(istr)
	errOrOk(err)
	log.Println(atoi)
	// 第二种字符串转数字的方式 字符串 进制 int类型
	log.Println("第二种字符串转数字的方式")
	parseInt, err2 := strconv.ParseInt("123", 10, 64)
	errOrOk(err2)
	log.Println(parseInt)
	// 数字转字符串
	log.Println("数字转字符串")
	itoa := strconv.Itoa(atoi)
	log.Println(itoa)

	log.Println("字符串转float32 float64 bool")
	// 字符串转float32 float64 bool
	// ParseFloat本质是无论后面写的是32还是64都是float64  只是按照32的格式转换
	fstr := "123.456"
	f1, ef1 := strconv.ParseFloat(fstr, 32)
	errOrOk(ef1)
	log.Println(f1)
	f2, ef2 := strconv.ParseFloat(fstr, 64)
	errOrOk(ef2)
	log.Println(f2)
	b1, bf1 := strconv.ParseBool(fstr)
	errOrOk(bf1)
	log.Println(b1)

	log.Println("基本类型转换为字符串")
	// 基本类型转换为字符串
	// bool转字符串
	formatBool := strconv.FormatBool(true)
	log.Println(formatBool)
	// int转字符串
	formatInt := strconv.FormatInt(123, 10)
	log.Println(formatInt)
	// float转字符串
	//f是格式
	//‘b’	-ddddp±ddd，二进制指数
	//‘e’	-d.dddde±dd，十进制指数
	//‘E’	-d.ddddE±dd，十进制指数
	//‘f’	-ddd.dddd，没有指数
	//‘g’	会根据实际情况选择合适的格式：-ddd.dddd 或 -d.dddde±dd
	//‘G’	同’g’，但使用大写的’E’
	//-1表示不限制小数点后的位数 大于的数字表示保留极为
	//32表示float32
	formartFloat := strconv.FormatFloat(123.456, 'f', -1, 32)
	log.Println(formartFloat)
}
func errOrOk(e error) {
	if e != nil {
		log.Println("转换错误")
	} else {
		log.Println("转换成功")
	}
}
