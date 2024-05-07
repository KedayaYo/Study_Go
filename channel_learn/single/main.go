package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}

}
func main() {
	// 默认情况下 channel是双向的
	// 但是，我们经常一个channel作为参数传递  希望对方是单向使用  不希望对方写数据只希望对方读数据
	//var ch1 chan int       // 双向
	//var ch2 chan<- float64 // 单向 只能写入
	//var ch3 <-chan string  // 单向 只能读取
	ch := make(chan int, 3)
	var send chan<- int = ch
	var recv <-chan int = ch
	send <- 1
	//<-send
	data := <-recv
	//data <- 1
	fmt.Println(data)

	// 无法将单向chan转换为双向chan  双向chan可以转换为单向chan
	//d1 := (chan int)(send)
	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(5 * time.Second)
}
