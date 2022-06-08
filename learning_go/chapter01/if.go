package main

import "fmt"

func main() {
	var yes string
	fmt.Scan(&yes)
	if yes == "y" || yes == "Y" {
		fmt.Println("西瓜")
	} else {
		fmt.Println("十个西瓜")
	}
}
