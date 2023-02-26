package main

import (
	"fmt"
)

func main() {
	// 常量为固定值, 后面不能进行赋值
	const NAME string = "lisi"
	fmt.Println(NAME)

	//省略类型
	// 定义多个产量，类型相同
	const a, b, c string = "AA", "BB", "CC"
	fmt.Println(a, b, c)

	// 定义多个产量，类型不相同
	const (
		d string = "DD"
		f int    = 11
	)
	fmt.Println(d, f)

	// 定义多个产量，省略类型
	const e, g = "EE", 12
	fmt.Println(e, g)

	const (
		C1 int = 1
		C2
		C3
	)
	fmt.Println(C1, C2, C3)

	// 枚举，const+iota
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1, E2, E3)

	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E4, E5, E6)

}
