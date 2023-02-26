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
	fmt.Println(s, a)

	// 变量类型通过值类型进行推导
	var ss, aa = "lisi", 23
	fmt.Println(ss, aa)

	// 这是一个简短声明，只能在函数内部使用
	isBoy := false
	fmt.Println(isBoy)

	/*
		变量只能声明一次，可以重新赋值
	*/
	ss, aa, isBoy = "zhangsan", 21 , true
	fmt.Println(ss, aa, isBoy)
	
	//赋值，s和ss相互赋值
	fmt.Println(s,ss)
	s, ss = ss, s
	fmt.Println(s,ss)
}
