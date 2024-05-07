package main

import "log"

func main() {
	a := 10
	log.Printf("a: %d\n", a>>1) // 5 除以2的n次方
	log.Printf("a: %d\n", a<<1) // 20 乘以2的n次方

	a >>= 2
	log.Printf("a: %d\n", a) // 10 相当于 a = a >> 2
	a <<= 2
	log.Printf("a: %d\n", a) // 40 相当于 a = a << 2

	var c *int
	b := &a // 取地址
	a -= 5
	c = b
	log.Printf("b: %d\n", *b)
	log.Printf("c: %d\n", *c)
}
