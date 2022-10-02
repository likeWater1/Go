package main

import (
	"fmt"
)

func totalizer() func(int) int {
	var n = 10
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	a := totalizer()
	fmt.Println(a(1))
}
