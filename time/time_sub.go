package main

import (
	"fmt"
	"time"
)

func main() {
	expiredTimeStr := "2022-02-17T11:31:19+08:00"
	loc, _ := time.LoadLocation("Local")
	expiredTime, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", expiredTimeStr, loc)
	if err == nil {
		theTime := time.Now()
		subTime := expiredTime.Sub(theTime)
		fmt.Printf("当前theTime:%v, expiredTime:%v, subTime:%v", theTime, expiredTime, subTime)
		if subTime.Hours() <= 15*24 {
			// lark通知
			fmt.Println("lark通知")
		} else {
			fmt.Printf("解析时间正常 expiredTime:%v", expiredTime)
		}
	} else {
		fmt.Println("解析时间异常")
		fmt.Printf("解析时间异常 err:%v", err)
	}
}
