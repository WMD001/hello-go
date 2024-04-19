package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if i := rand.Intn(10); i > 5 {
		fmt.Println(i, "> 5")
	} else if i > 3 {
		fmt.Println(i, "> 3")
	} else {
		fmt.Println(i, "<= 3")
	}
}
