package main

import (
	"fmt"
	"time"
)

func go_worker(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("启动一个协程", name, "---")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, " 执行完毕")
}

func main() {
	go go_worker("第一个")
	go go_worker("第二个")

	for {
		time.Sleep(1 * time.Second)
	}
}
