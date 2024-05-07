package main

import "fmt"

func main() {
	// for循环
	/*
		for  {

			}
		等同于while(true)
	*/

	a := "Kedaya浩浩"
	ars := []rune(a)
	for i := 0; i < len(ars); i++ {
		fmt.Printf("下标: %d", i)
		fmt.Printf(", 值: %s\n", string(ars[i]))
	}

	// for range 主要针对 字符串、数组、切片、map、channel
	/*
		for key,value := range a {

		}
	*/
	for k, v := range ars {
		//if "a" == string(v) {
		//continue
		//break
		//}
		fmt.Printf("下标: %d", k)
		fmt.Printf(", 值: %s\n", string(v))
	}
}
