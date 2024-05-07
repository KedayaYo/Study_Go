package main

import "fmt"

type MyReader interface {
	Read(content string)
}
type MyWriter interface {
	Write(content string)
}
type MyReadWriter interface {
	MyReader
	MyWriter
	ReadWrite()
}

// 具体实现
type ReadWriter struct {
}

func (r *ReadWriter) Read(content string) {
	fmt.Println("Read: ", content)
}

func (r *ReadWriter) Write(content string) {
	fmt.Println("Write: ", content)
}

func (r *ReadWriter) ReadWrite() {
	fmt.Println("ReadWrite")
}

// 接口的嵌套
func main() {
	// 第一种方式 r ReadWriter适用于实现r ReadWriter 的结构体 而不是指针类型  这种时候  ReadWriter{}和&ReadWriter{}都可以  但是这个时候是值传递 你得保证不影响到正常逻辑
	// 第二种方式 r *ReadWriter只能适用于实现r *ReadWriter 的结构体  这种时候只能是&ReadWriter{}  这个时候是指针传递
	//var rw MyReadWriter = ReadWriter{}
	var rw MyReadWriter = &ReadWriter{}
	rw.Read("hello")
}
