package main

/*
在声明一个变量而不指定其类型时（即使用不带类型的 := 语法 var = 表达式语法），变量的类型会通过右值推断出来。

当声明的右值确定了类型时，新变量的类型与其相同
*/

import "fmt"

func main() {
	v := 42 // 修改这里看看！
	fmt.Printf("v is of type %T\n", v)
}
