package main

import (
	"fmt"
)

func main() {
	s := ""
	b := &s
	fmt.Println("打印", *b == "")
}
