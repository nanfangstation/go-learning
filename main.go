package main

import (
	"fmt"
	"net/url"
)

func main() {
	key := "xxx"
	key = url.QueryEscape(key)
	fmt.Print(key)
}
