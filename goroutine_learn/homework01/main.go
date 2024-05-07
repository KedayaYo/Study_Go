package main

import (
	"fmt"
	"time"
)

// 标识
var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		// 接收 接收到number有true输入
		<-number
		// 让另一个携程来通知该数字打印  也就是使用无缓冲channel
		fmt.Printf("%d%d", i, i+1)
		// 打印i和i+1  因此每次加2
		i += 2
		letter <- true
	}
}

func printLetter() {
	// 此处是下标
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		// 接收 接收到letter有true输入
		<-letter
		if i >= len(str) {
			return
		}
		// 让另一个携程来通知该数字打印  也就是使用无缓冲channel
		fmt.Printf(str[i : i+2]) // 左闭右开
		// 打印i和i+1  因此每次加2
		i += 2
		number <- true
	}
}
func main() {
	// 需求：使用两个携程 一个打印数字 一个打印字母 交替打印 实现效果如下：12AB34CD56EF78GH910IJ
	go printNum()
	go printLetter()
	number <- true

	time.Sleep(5 * time.Second)
}
