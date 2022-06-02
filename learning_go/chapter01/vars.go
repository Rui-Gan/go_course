package main

import "fmt"

var vesion string = "1.0"

func main() {
	var me string
	fmt.Println(me)
	me = "kk"
	fmt.Println(me)

	var name, user string = "kk", "woniu"
	fmt.Println(name, user)

	//定义多个不同类型的变量

	var (
		age    int     = 21
		height float64 = 1.63
	)
	fmt.Println(age, height)

	var (
		s = "kk"
		a = 31
	)
	fmt.Println(s, a)
	var ss, aa = "ll", 22
	fmt.Println(ss, aa)

	//简短声明 函数内部使用
	isBoy := true
	fmt.Println(isBoy)
	s, ss = ss, s
	fmt.Println(s, ss)

}
