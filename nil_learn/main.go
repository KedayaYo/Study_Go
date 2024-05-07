package main

func main() {
	// Go 不同数据零值不一样
	/*
		bool: false
		int: 0
		float: 0
		string: ""
		pointer: nil
		slice: nil
		map: nil
		channel: nil
		interface: nil
		function: nil
		struct 默认值不是nil 根据具体字段类型的默认值来确定  判断相等也是每个字段都要相等
	*/
	// 例如可以用make初始化的数据类型：slice、map、channel 即使设置为make([]Person, 0) 也不是nil 是个空的slice
	// 即使slice本质是一个struct 但是里面有三个字段：指向数组的指针、长度、容量  判断的是指针的零值  所以var a []int 是等于nil的

}
