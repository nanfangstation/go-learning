package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Print(a)
	var b [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = i + j
		}
	}
	fmt.Print("b:", b)
}
