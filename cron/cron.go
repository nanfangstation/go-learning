package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	i := 0
	c := newWithSeconds()
	spec := "*/1 * * * * ?" //一分钟运行一次
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()

	select {}
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
