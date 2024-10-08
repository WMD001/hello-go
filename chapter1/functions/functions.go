package main

/*
函数可接受零个或多个参数。

在本例中，add 接受两个 int 类型的参数。

注意类型在变量名的 后面。

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

*/

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
}
