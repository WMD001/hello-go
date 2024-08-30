package main

/*
推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。
*/

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, ",")
	}
}

// 输出 9,8,7,6,5,4,3,2,1,0,
