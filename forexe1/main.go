package main

import (
	"fmt"
)

/*

  *
 ***
*****

*每一行是2的倍数-1，2*n-1
空格是总层数减当前层数


    *
   * *
  *   *
 *     *
*********
*/
func main() {
	total := 5
	for i := 1; i <= total; i++ {
		for k := 1; k <= total-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == total {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

	//i表示层数
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}
