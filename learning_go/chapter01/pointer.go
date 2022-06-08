package main

import "fmt"

func main() {
	a := 2
	b := a
	b = 3
	fmt.Println(a, b)

	aa := &a
	bb := aa
	*bb = 3
	fmt.Println(*aa, *bb, a)

}
