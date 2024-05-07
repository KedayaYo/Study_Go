/**
 * @Author Kedaya
 * @Date 2024/5/5 15:47:00
 * @Desc
 **/
package user

import "fmt"

func init() {
	fmt.Println("user init")
}

// package 用来组织源码  是多个go源码的集合 代码复用的基础 fmy os io
// 同一个文件夹下 一个包只能有一个package
type User struct {
	Name string
	gen  string
}
