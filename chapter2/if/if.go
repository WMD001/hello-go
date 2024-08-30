package main

/*
Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( )，而大括号 { } 则是必须的。
*/

import "fmt"

func check(i int) string {
	if i > 0 {
		return "true"
	}
	return "false"
}

func main() {
	fmt.Println(check(1))
}
