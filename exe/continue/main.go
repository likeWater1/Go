package main

import "fmt"

func main() {
	//输出100内的奇数
	/*
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				continue
			}
			fmt.Printf("i=%d\n", i)

		}
	*/
	/*100,0000元，经过一次路口缴费
	1、现金大于5,000，缴费5%
	2、现金小于等于5,000，缴费1000
	计算经过多少个路口
	*/
	var money float32 = 100000
	var sum int
	for {
		if money < 1000 {
			break
		}
		if money > 50000 {
			money -= money * 0.05
		} else {
			money -= 1000
		}

		sum++

	}
	fmt.Printf("经过%v个路口，剩余%v元", sum, money)
}
