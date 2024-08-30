package main

/*
在 if 的简短语句中声明的变量同样可以在对应的任何 else 块中使用。
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	if i := rand.Intn(10); i > 5 {
		fmt.Println(i, "> 5")
	} else if i > 3 {
		fmt.Println(i, "> 3")
	} else {
		fmt.Println(i, "<= 3")
	}
}
