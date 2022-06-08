package main

import "fmt"

func main() {
	var op string
	fmt.Scan(&op)

	if op != "y" {
		goto END
	}
	fmt.Println("1")

END:
	i := 1
	result := 0
START:
	if i > 100 {
		goto NEXT
	} else {
		result += i
		i++
		goto START
	}
NEXT:
	fmt.Println(result)

BREAKPOINT:
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i*j == 9 {
				break BREAKPOINT
			} else {
				fmt.Println(i * j)
			}
		}
	}
}
