package main

import (
	"fmt"
)

func main() {
	//函数类型。
	type fun func(int, int) int
	var ff fun = add
	fmt.Println("ff", ff(2, 3))
	var ff1 fun = sub
	fmt.Println("ff1", ff1(4, 2))
	name, age := b()
	fmt.Println("name=", name)
	fmt.Println("age=", age)

	n2 := 2
	f1(n2)
	fmt.Println("d: \n", n2)

}
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

//函数返回值

func b() (name, age string) {
	name = "tom"
	age = "10"
	return
}

func f1(x int) {
	x = 1
}
