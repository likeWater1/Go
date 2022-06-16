package main

import (
	"fmt"
)

func main() {
	h := "1"
	switch h {
	case "1":
		fmt.Println("111")
		fallthrough //穿透，fallthrough 会强制执行后面的 case 语句。
	case "2":
		fmt.Println("222")
	default:
		fmt.Println("defaulet")
	}

	// fmt.Println("输入")
	// var b string
	// fmt.Scanln(&b)
	// switch {
	// case b == "a":
	// 	fmt.Println("a")
	// case b == "b":
	// 	fmt.Println("b")
	// }

	var x interface{}
	var y int = 123
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Println("空")
	case int:
		fmt.Printf("int:%T\n", i)
	case float32:
		fmt.Println("float")
	}

	//小练习
	fmt.Println("输入分数")
	var char float64
	fmt.Scanln(&char)
	switch {
	case char >= 60.0:
		fmt.Println("及格")
	case char < 60.0:
		fmt.Println("不及格")
	}
}
