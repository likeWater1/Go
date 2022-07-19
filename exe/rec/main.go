package main

import (
	"fmt"
)

func rec(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*rec(n-1) + 1
	}
}

func main() {

	fmt.Println("rec(3)=", rec(3))
}

//等于7。解析：n=2时，进行else。2 * rec (2-1) + 1。又一次调用rec函数。rec (2-1)=rec（1）。n=1，返回3。2*3+1=7
