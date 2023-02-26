package main

import (
	"fmt"
)

var version string = "2.0"

func main() {
	var me string
	me = "lisp"
	fmt.Println(me)
	fmt.Println(version)

	//定义相同类型
	var name, user string = "lisi", "zhangsan"
	fmt.Println(name, user)

	// 定义不同类型变量
	var (
		age    int     = 23
		height float64 = 1.22
	)
	fmt.Println(age, height)

	var (
		s = "lisi"
		a = 23
	)
	fmt.Println(s,a)

	var ss , aa = "lisi" , 23
	fmt.Println(ss,aa)
}
