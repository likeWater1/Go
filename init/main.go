package main

import (
	"fmt"
	"src/go_code/project01/init2"
)

//init函数会在main函数前调用。全局变量会在init函数前调用。  全局变量 > init() > main()
//init函数主要作用是，完成一些初始化操作
//

var age = test()

func test() int {
	fmt.Println("test()")
	return 10
}

func init() {
	fmt.Println("init()")
}

func main() {
	fmt.Println("main(),age=", age)
	fmt.Println("Num=", init2.Num)
}
