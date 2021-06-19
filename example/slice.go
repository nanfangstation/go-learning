package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	two := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		two[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			two[i][j] = i + j
		}
	}
	print(two)
}
