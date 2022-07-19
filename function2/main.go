//函数细节
package main

import (
	"fmt"
)

//函数也是一种数据类型，可赋值给变量，则该变量就是函数类型的变量。通过变量可以对该函数进行调用
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

//函数充当形参
func myFun(funvar myFuncVar, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//数据类型别名，type，函数类型重命名
type myFuncVar func(n1 int, n2 int) int

//返回值重命名
func ret(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//支持多个参数
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0]表示取出args切片的第一个元素值。
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a的数据类型:%T\n", a)

	res := a(10, 20) //相当于 res := getSum(10,20)
	fmt.Printf("res=%d\n", res)

	//函数充当行参
	fun := myFun(getSum, 20, 30)
	fmt.Println("fun=", fun)

	//数据类型别名
	type myInt int
	var num4 int
	var num3 myInt = 10
	fmt.Printf("num3的数据类型是:%T,num3=%d\n", num3, num3) //数据类型main.myInt
	//num4 = num3 虽然对对int类型进行别名，但go认为myInt和int仍然是两个类型。需要类型转换
	num4 = int(num3)
	fmt.Printf("num4的数据类型是:%T,num4=%d\n", num4, num4)

	//函数类型重命名
	c := myFun(getSum, 10, 20)
	fmt.Println("c=", c)

	//返回值重命名
	a1, b1 := ret(1, 1)
	fmt.Printf("a=%d,b=%d\n", a1, b1)

	//支持多个参数
	d := sum(10, 10)
	fmt.Print("d=", d)
}
