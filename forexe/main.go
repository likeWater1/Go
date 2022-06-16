package main

import (
	"fmt"
)

func main() {
	//3个班五个学生，每人的成绩以及每个班平均分

	/*	var studentNum int8 = 3
		var classNum int8 = 3
		//var count float64   错误记录，写在这里的话，总算不会初始化，比如，第一个班级三位成绩都是10，第一次允许完后 count=30，平均成绩30/3=10
		//第二个班级三位成绩都是10，count +=cj等于 count = count + cj也就是count = 30+ 30 = 60。把一班的总成绩也加进来啦。
		//所以要将初始化的变量放入循环体内，每次进行初始化
		var total int8
		for class := 1; class <= int(classNum); class++ {
			var count float64
			for i := 1; i <= int(studentNum); i++ {
				var cj float64
				fmt.Printf("请输入班级%v,第%d学生的成绩:\n", class, i)
				fmt.Scanln(&cj)
				count += cj
			}
			fmt.Printf("班级%v平均成绩为%v\n", class, count/float64(studentNum))
			total += int8(count)
		}
		fmt.Printf("全部班级总成绩：%v\n", total)

		/*var total int
		for class := 1; class <= 3; class++ {
			var count int
			for i := 1; i <= 3; i++ {
				var cj int
				// var count int 写在这里的话，每次循环会初始化为0，count = 0 + 10。所以要写在循环外。
				fmt.Printf("请输入第%v位同学的成绩\n", i)
				fmt.Scanln(&cj)
				//fmt.Printf("第%v位的成绩:%v\n", i, cj)
				count += cj
			}
			fmt.Printf("第%v班级平均成绩:%v\n", class, count/3)
			total += count
		}
		fmt.Printf("全部班级总成绩：%v", total)
	*/
	var b int
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			if b%2 == 0 {
				fmt.Printf("%d\n", b)
				break
			}
			fmt.Printf("%d\n", b)
		}
		fmt.Printf("%d\n", b)
	}
}
