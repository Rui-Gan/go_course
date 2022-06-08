package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Problem 1:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
	fmt.Println("Problem 2:")
GAME:
	println("请输入一个数，您最多有五次机会：")
	rand.Seed(time.Now().Unix())
	num := rand.Int() % 10
	fmt.Println(num)
	for cnt := 5; cnt >= 0; cnt-- {
		if cnt == 0 {
			fmt.Println("太笨了，您将重新开始游戏")
			goto GAME
		}
		var a int
		fmt.Scan(&a)
		fmt.Println(a)
		if a == num {
			fmt.Println("恭喜您猜中了")
			break
		} else if a > num {
			fmt.Println("太大了，您还有", cnt-1, "次机会")
		} else {
			fmt.Println("太小了，您还有", cnt-1, "次机会")
		}
	}
	fmt.Println("退出游戏")
}
