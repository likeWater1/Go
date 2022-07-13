package main

import (
	_ "crypto/rand"
	"fmt"
	_ "math/rand"
	_ "time"
)

func main() {
	/*
		var num int
		for {
			rand.Seed(time.Now().Unix())
			i := rand.Intn(5)
			fmt.Println(i)
			num++
			if num == 6 {
				break
			}
		}
		*/
labels1: //跳出整个循环
	for i := 0; i < 4; i++ {
		//labels1:跳出第二个循环
		for j := 0; j < 3; j++ {
			if j == 2 {
				break labels1
			}
			fmt.Println(j)
		}
	}

	//输出三次密码，第三次错误跳出循环

	for i := 1; i <= 3; i++ {
		var password string
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if password == "321" {
			break
		}
	}
}
