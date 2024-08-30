package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, ",")
	}
}

// 输出 9,8,7,6,5,4,3,2,1,0,
