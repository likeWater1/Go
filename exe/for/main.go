package main

import (
	_ "crypto/rand"
	"fmt"
	_ "math/rand"
	_ "time"
)

func main() {
	/*
		var num int
		for {
			rand.Seed(time.Now().Unix())
			i := rand.Intn(5)
			fmt.Println(i)
			num++
			if num == 6 {
				break
			}
		}
	*/
labels1:
	for i := 0; i < 4; i++ {
		//labels1:
		for j := 0; j < 3; j++ {
			if j == 2 {
				break labels1
			}
			fmt.Println(j)
		}
	}
}

