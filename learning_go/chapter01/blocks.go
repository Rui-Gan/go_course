package main

import "fmt"

func main() {
	//作用域：定义标识符可以使用的范围
	//用{}来定义作用域的范围
	outer := 1

	{
		inner := 2
		fmt.Println(inner)
		fmt.Println(outer)
	}
	//fmt.Println(inner)

}
