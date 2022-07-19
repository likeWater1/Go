package main

import (
	"fmt"
)

func swap(n1, n2 *int) {
	//定义个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("n1=%d,n2=%d", a, b)
}
