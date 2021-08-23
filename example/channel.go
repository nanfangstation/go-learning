package main

import (
	"fmt"
	"reflect"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("发送前: ", c, reflect.TypeOf(c))
	c <- sum // 把sum发送到通道c
	fmt.Println("发送后: ", c, reflect.TypeOf(c))
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
