package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	i := 0
	z := float64(1)
	for ; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(float64(2)))
	fmt.Println(math.Sqrt(float64(2)))
}
