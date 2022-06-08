package main

import "fmt"

func main() {
	result := 0
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)
	result = 0
	i := 1
	for i <= 100 {
		result += i
		i++
	}
	fmt.Println(result)
	i = 0
	for {
		if i == 1 {
			break
		}
		fmt.Println(i)
		i++
	}

	//字符串，数组，切片，映射，管道
	desc := "wach"
	for i, ch := range desc {
		fmt.Printf("%T %T\n", i, ch)
	}
}
