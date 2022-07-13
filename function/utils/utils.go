package utils

import "fmt"

//函数名首字母必须大写。才可以引用。改函数可导出
func Cal(n1 float64, n2 float64, opt byte) float64 {
	var res float64
	switch opt {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	default:
		fmt.Print("操作失误")
	}
	//fmt.Print(res)
	return res
}
