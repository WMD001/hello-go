package main

/*
和 for 一样，if 语句可以在条件表达式前执行一个简短语句。

该语句声明的变量作用域仅在 if 之内。

*/

import (
	"fmt"
	"math/rand"
)

func main() {
	if i := rand.Intn(10); i > 5 {
		fmt.Println("i is", i)
	}
	fmt.Println("over")
}
