package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"
	"time"
)

//声明接口类
type EchoService struct{}

//定义方法Echo
func (service *EchoService) Echo(arg string, result *string) error {
	*result = arg
	return nil
}

//服务端启动逻辑
func RegisterAndServeOnTcp() {
	err := rpc.Register(&EchoService{}) //注册并不是注册方法，而是注册EchoService的一个实例
	if err != nil {
		log.Fatal("error registering", err)
		return
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("error resolving tcp", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error accepting", err)
		} else {
			//这里先通过NewServerCodec获得一个实例，然后调用rpc.ServeCodec来启动服务
			rpc.ServeCodec(msgpk.NewServerCodec(conn))
		}
	}
}

//客户端调用逻辑
func Echo(arg string) (result string, err error) {
	var client *rpc.Client
	conn, err := net.Dial("tcp", ":1234")
	client = rpc.NewClientWithCodec(msgpk.NewClientCodec(conn))

	defer client.Close()

	if err != nil {
		return "", err
	}
	err = client.Call("EchoService.Echo", arg, &result) //通过类型加方法名指定要调用的方法
	if err != nil {
		return "", err
	}
	return result, err
}

//main函数
func main() {
	go server.RegisterAndServeOnTcp() //先启动服务端
	time.Sleep(1e9)
	wg := new(sync.WaitGroup) //waitGroup用于阻塞主线程防止提前退出
	callTimes := 10
	wg.Add(callTimes)
	for i := 0; i < callTimes; i++ {
		go func() {
			//使用hello world加一个随机数作为参数
			argString := "hello world " + strconv.Itoa(rand.Int())
			resultString, err := client.Echo(argString)
			if err != nil {
				log.Fatal("error calling:", err)
			}
			if resultString != argString {
				fmt.Println("error")
			} else {
				fmt.Printf("echo:%s\n", resultString)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
