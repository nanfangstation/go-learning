package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [5]int
	fmt.Println(a)
	var b [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = i + j
		}
	}
	fmt.Println("b:", b)
	c := [...]string{1: "1", 2: "3", 4: "5"}
	fmt.Println("c:", c, len(c), reflect.TypeOf(c))
}
