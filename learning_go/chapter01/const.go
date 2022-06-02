package main

import "fmt"

func main() {
	const name string = "kk"
	fmt.Println(name)

	//省略类型
	//定义多个常量（类型相同）
	//定义多个常量（类型不相同）
	//定义多个常量（省略类型）
	const age = 10
	const a, b, c string = "a", "b", "c"
	const (
		aa string  = "aa"
		bb int     = 18
		cc float64 = 1.98
	)
	const c1, c2 = "sile", 1
	fmt.Println(age, a, b, c, aa, bb, cc, c1, c2)

	//常量和前面一样，可以直接写标识符
	const (
		c3 int = 1
		c4
		c5
		c6 float64 = 3.14
		c7
		c8
		c9 string = "kk"
	)
	fmt.Println(c3, c4, c5, c6, c7, c8, c9)
	//枚举 const+iota
	const (
		e1 int = iota*3 + 5
		e2
		e3
	)
	fmt.Println(e1, e2, e3)

}
