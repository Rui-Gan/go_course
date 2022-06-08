package main

import "fmt"

func main() {
	var name string = `k\tk` //原生字符串
	var desc string = "w\tw" //可解释字符串

	fmt.Println(name, desc)
}
