package main

import (
	"fmt"
)

func main() {
	// var time float64
	// var gender string
	// fmt.Println("成绩")
	// fmt.Scanln(&time)
	// if time <= 8 {
	// 	fmt.Println("性别")
	// 	fmt.Scanln(&gender)
	// 	if gender == "男" {
	// 		fmt.Println("进入男子组")
	// 	} else if gender == "女" {
	// 		fmt.Println("进入女子组")
	// 	} else if gender != "女" && gender != "男" {
	// 		fmt.Println("输入正确性别")
	// 	}
	// } else {
	// 	fmt.Println("淘汰")
	// }

	////案例2
	var month byte
	var age byte
	fmt.Println("月份")
	fmt.Scanln(&month)
	fmt.Println("年龄")
	fmt.Scanln(&age)
	if month > 4 && month < 10 {
		if age < 18 {
			fmt.Println("儿童半价")
		} else if age > 60 {
			fmt.Println("1/3")
		} else {
			fmt.Println("成人60")
		}
	} else {
		if age > 18 && age < 60 {
			fmt.Println("成人40")
		} else {
			fmt.Println("其他40")
		}
	}
}
