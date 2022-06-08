package main

import "fmt"

func main() {
	var yes string
	fmt.Scan(&yes)

	switch yes {
	case "y", "n":
		fmt.Println("1")
	case "Y":
		fmt.Println("2")
	}
	//fmt.Printf("成绩")
	var score int
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
