package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	fmt.Print(m)
	_, prs := m["2"]
	fmt.Println("prs:", prs)
}
