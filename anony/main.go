package main

import (
	"fmt"
)

//将匿名函数赋值给一个全局变量。全局匿名函数
var (
	//func1  就是个全局匿名函数
	func1 = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {

	//定义匿名函数直接调用，这种函数只能调用一次
	//演示  匿名函数不写函数名。（10，20）是调用此函数。
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	//将匿名函数赋值给a。然后再调用。a的数据类型就是函数类型
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)
	fmt.Println(res2)

	//全局匿名函数
	res3 := func1(40, 1)
	fmt.Println(res3)

}
