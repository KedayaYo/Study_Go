package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex
var wp sync.WaitGroup

func Write(num *int) {
	time.Sleep(3 * time.Second)
	defer wp.Done() // 确保在函数退出时调用 Done
	rwLock.Lock()
	defer rwLock.Unlock()
	*num += 1
	fmt.Println("write num:", *num)
	time.Sleep(3 * time.Second)
}
func Read(num *int) {
	defer wp.Done() // 确保在函数退出时调用 Done
	for {
		rwLock.RLock()
		//defer rwLock.RUnlock() // 这是死循环  永远走不到这里
		time.Sleep(500 * time.Millisecond)
		fmt.Println("read num:", *num)
		rwLock.RUnlock()
	}
}

// 读写锁
func main() {

	num := 10

	wp.Add(2)
	// 读锁
	go Read(&num)
	// 写锁
	go Write(&num)
	wp.Wait()
	fmt.Printf("main end：%d\n", num)
}
