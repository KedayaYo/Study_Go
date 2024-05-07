package main

import (
	"fmt"
	_ "go_learn/package_learn/user" // 可以使用匿名  这种方式只会执行包的init方法  不会执行包的其他方法的时候使用
	// . "go_learn/package_learn/user" //可以使用.  可以直接访问User GetUsername 但是不推荐 会导致代码可读性降低不知道是本地方法还是引入的方法
	u "go_learn/package_learn/user" //可以使用别名  这样只能通过别名访问
)

func main() {
	// 使用包名.变量名
	user := u.User{
		Name: "Entic",
	}
	user.Name = "Kedaya"
	// 不同包 需要大写表示public
	//user.gen = "男"

	username := u.GetUsername(user)
	fmt.Println(username)
}
