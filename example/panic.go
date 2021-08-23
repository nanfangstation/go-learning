package main

import (
	"os"
	"runtime"
)
import "fmt"

var user = os.Getenv("USER")

func main() {
	defer printStack()
	fmt.Println("user:", user)
	//panic("no value for $USER")
	if user == "" {
		panic("no value for $USER")
	}
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
