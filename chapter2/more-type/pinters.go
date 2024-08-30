package main

import "fmt"

func main() {
	var p *int
	i := 42

	p = &i

	fmt.Println(p)  //0xc00000a0a8
	fmt.Println(i)  //42
	fmt.Println(*p) //42

}
