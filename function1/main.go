package main

import (
	"fmt"
)

//。没有reture情况下。函数在栈中是独立的
//基本数据类型和数组默认是值传递，即指拷贝，在函数内修改不会影响，原来的值。
func test(n1 int) {
	n1 += 1
	fmt.Printf("test()n1=%d\n", n1)
}

//如果希望函数内的变量能修改函数外的变量，可以通过传入变量的内存地址&，函数内以指针的方式操作变量。类似引用
func test01(n2 *int) {
	*n2 += 1
	fmt.Printf("n2的地址:%v\n", &n2)
	fmt.Printf("n2=%v\n", *n2)
}

//return，将函数执行结果返回给调用者
func sum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Printf("sum():%d\n", sum)
	return sum
}

//返回多个值
func getSumSub(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

//递归调用
func rec(n1 int) {
	if n1 > 2 {
		n1--
		rec(n1)
	}
	fmt.Printf("n1=%d\n", n1)
}

//主
func main() {
	n1 := 10
	test(n1)
	fmt.Printf("main()n1=%d\n", n1)

	//如果希望函数内的变量能修改函数外的变量，可以通过传入变量的内存地址&，函数内以指针的方式操作变量。类似引用
	num := 20
	test01(&num)
	fmt.Printf("num的地址=%v\n", &num)
	fmt.Printf("main()num=%v\n", num)
	//return，将函数执行结果返回给调用者
	sum := sum(10, 20)
	fmt.Printf("sum main():%d\n", sum)

	//返回多个值
	reSum, reSub := getSumSub(1, 2)
	fmt.Printf("sum:%d,sub:%d\n", reSum, reSub)

	//递归调用
	rec(4)
}
