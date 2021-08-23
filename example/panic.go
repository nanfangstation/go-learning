package main

import "os"
import "fmt"

var user = os.Getenv("USER")

func main() {
	fmt.Println("user:", user)
	panic("no value for $USER")
	if user == "" {
		panic("no value for $USER")
	}
}
