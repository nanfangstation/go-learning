package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	//listener, err := net.Listen("tcp", "localhost:8080")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Print(err)
	//	}
	//	go handleConn(conn)
	//}
	//key := "xxx.xlsx"
	//fmt.Println(key)
	//key = url.QueryEscape(key)
	//fmt.Println(key)
	s := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int, 3)
	var successCount = 0
	var failCount = 0
	fmt.Println(successCount, failCount)

	go sum(s, ch, successCount, failCount)
	successCount, failCount = <-ch, <-ch // 从通道 c 中接收

	fmt.Println(successCount, failCount)
	beginTimeStr := time.Unix(1653276180, 0).Format("2006-01-02 15:04:05")
	fmt.Printf(beginTimeStr)
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // 连接断开
		}
		time.Sleep(1 * time.Second)
	}
}

func sum(s []int, c chan int, successCount, failCount int) {
	c <- successCount // 把 sum 发送到通道 c
	c <- failCount    // 把 sum 发送到通道 c
	sum := 0
	for _, v := range s {
		successCount = successCount + 1
		sum += v
	}
}
