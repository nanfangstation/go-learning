package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS)
	//wg.Add(1)
	go func() {
		fmt.Println(runtime.NumGoroutine())
		log.Println(runtime.NumGoroutine())
		time.Sleep(time.Second * 2)
		wg.Wait() // 阻塞在此
	}()
	//wg.Wait() // 阻塞在此
}
