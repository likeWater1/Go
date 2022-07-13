package main

import (
	"fmt"
	"project01/function/utils/utils"
)

var sum float64

func main() {
	var n1 float64 = 1
	var n2 float64 = 2
	a := utils.Cal(n1, n2, opt)
	fmt.Println(a)
}
