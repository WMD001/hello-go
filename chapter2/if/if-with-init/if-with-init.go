package main

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
