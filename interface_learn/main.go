package main

import "fmt"

// 定义接口
type Duck interface {
	// 方法声明
	Gaga()
	Walk()
	Swimming()
}

// 接口实现
type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎嘎")
}

func (pd *pskDuck) Walk() {
	fmt.Println("走路")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("游泳")
}

type MyWriter interface {
	Write(content string) error
}

type MyReader interface {
	Reader(content string) error
}

type writerAndReader struct {
	// 接口也是类型  也可以声明在结构体中
	MyWriter
}

//func (wr *writerAndReader) Write(content string) error {
//	fmt.Println("write content: ", content)
//	return nil
//}
//
//func (wr *writerAndReader) Reader(content string) error {
//	fmt.Println("read content: ", content)
//	return nil
//}

// 验证interface在结构体中的使用
type fileWriter struct {
	filePath string
}

func (fw *fileWriter) Write(content string) error {
	fmt.Println("file write content: ", content)
	return nil
}

type databaseWriter struct {
	host string
	port int
	db   string
}

func (dw *databaseWriter) Write(content string) error {
	fmt.Println("database write content: ", content)
	return nil

}

func main() {
	// 鸭子类型
	// Go语言处处都是interface 处处都是鸭子类型
	// 鸭子类型强调的是外部行为，而不是内部结构

	// 接口实现使用
	// func (pd *pskDuck) Swimming() 实现方法的时候 入参是指针类型  需要实现结构体的指针类型
	//var duck Duck = pskDuck{} // Type does not implement Duck as the Gaga method has a pointer receiver
	var duck Duck = &pskDuck{}
	duck.Gaga()
	duck.Walk()
	duck.Swimming()

	// 多接口实现  这是一个多对多的关系 一个接口可以被多个结构体实现 一个结构体也可以实现多个接口
	//var wr MyWriter = &writerAndReader{}
	//var rd MyReader = &writerAndReader{}
	//wr.Write("hello")
	//rd.Reader("World")

	// 接口在结构体中的使用
	// 好处就是  在可以通过接口的方式来调用各种结构体  无论结构体的本身有多少字段 都可以轻易拿到
	// 这个时候writerAndReader就不能自己实现WyWriter了 需要通过fileWriter和databaseWriter来实现
	var fw MyWriter = &writerAndReader{
		&databaseWriter{},
	}
	fw.Write("file write content")
}
