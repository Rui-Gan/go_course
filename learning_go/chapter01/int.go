package main

import "fmt"

func main() {
	var age int = 31
	fmt.Printf("%d\n", age)

	//位运算 & ｜ ^(异或) &^(按位清空) >> <<

	//int unit byte rune int* 类型转换
	var c int = 0xfc11
	fmt.Println(uint8(c))

	//byte rune
	var a byte = 'a'
	var w rune = '中'
	fmt.Printf("%x\n", age)
	fmt.Println(a, w)
	fmt.Printf("%T %q %U", w, w, w)
}
