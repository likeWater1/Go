package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string转化基本类型
	var a string = "123"
	var b int64
	b, _ = strconv.ParseInt(a, 10, 64)
	fmt.Printf("b的类型%T,b=%v\n", b, b)
	//基本转string
	var c int64 = 1
	//var d string
	d := fmt.Sprintf("%d", c)
	fmt.Printf("d的类型%T,b=%v", d, d)
}
