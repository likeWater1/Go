package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	//练习
	k := 0
	for {
		if k <= 10 {
			fmt.Println(k)
		} else {
			break
		}
		k++
	}
	//len() 输出字符串长度
	var str string = "abcd"
	fmt.Println("str的长度：", len(str))

	//练习2
	//遍历1
	//传统遍历方式。若字符串中含有中文，会出现乱码。因为传统遍历是通过字节的方式进行遍历。一个中文站3字节
	var str1 string = "abcdefd"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i]) //str1[i]显示的是字符对应的码值，需要通过%c转换
	}
	//
	//遍历2
	//for-range遍历。若字符串中含有中文，则不会出现问题。因为是按照字符的形式进行遍历。
	for index, val := range str1 {
		fmt.Printf("index:%v,val:%c\n", index, val)
	}

	//练习3
	//1～100中9的倍数的整数个数及总和
	var sum int
	var count uint
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("总和：%d，个数：%d", sum, count)
}
