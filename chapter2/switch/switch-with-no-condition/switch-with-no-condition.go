package main

/*
无条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰。
*/

import "time"

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 19:
		println("Good afternoon.")
	default:
		println("Good evening.")
	}
}
