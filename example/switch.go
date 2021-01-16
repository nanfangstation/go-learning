package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("tow")
	default:
		fmt.Print("default")
	}
}
