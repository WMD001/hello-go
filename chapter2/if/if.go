package main

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
