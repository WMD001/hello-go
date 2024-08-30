package main

/*
初始化语句和后置语句是可选的。你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
*/

import "fmt"

func main() {
	i := 0
	for i < 10 {
		i++
	}
	fmt.Println(i)
}
