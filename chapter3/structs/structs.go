package main

/*
一个 结构体（struct）就是一组 字段（field）。
*/

import "fmt"

// Vertex 定义结构体
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
}
