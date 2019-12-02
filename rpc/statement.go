package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
	"net/rpc"
)

// 定义请求/响应的完整数据
type MsgRequest struct {
	rpc.Request             //head
	Arg         interface{} //body
}

type MsgResponse struct {
	rpc.Response             //head
	Reply        interface{} //body
}

// 自定义Codec的声明
type MessagePackServerCodec struct {
	rwc    io.ReadWriteCloser //用于读写数据，实际是一个网络连接
	req    MsgRequest         //用于缓存解析到的请求
	closed bool               //标识codec是否关闭
}

type MessagePackClientCodec struct {
	rwc    io.ReadWriteCloser
	resp   MsgResponse //用于缓存解析到的请求
	closed bool
}

func NewServerCodec(conn net.Conn) *MessagePackServerCodec {
	return &MessagePackServerCodec{conn, MsgRequest{}, false}
}

func NewClientCodec(conn net.Conn) *MessagePackClientCodec {
	return &MessagePackClientCodec{conn, MsgResponse{}, false}
}

func readData(conn io.ReadWriteCloser) (data []byte, returnError error) {
	const HeadSize = 4 //设定长度部分占4个字节
	headBuf := bytes.NewBuffer(make([]byte, 0, HeadSize))
	headData := make([]byte, HeadSize)
	for {
		readSize, err := conn.Read(headData)
		if err != nil {
			returnError = err
			return
		}
		headBuf.Write(headData[0:readSize])
		if headBuf.Len() == HeadSize {
			break
		} else {
			headData = make([]byte, HeadSize-readSize)
		}
	}
	bodyLen := int(binary.BigEndian.Uint32(headBuf.Bytes()))
	bodyBuf := bytes.NewBuffer(make([]byte, 0, bodyLen))
	bodyData := make([]byte, bodyLen)
	for {
		readSize, err := conn.Read(bodyData)
		if err != nil {
			returnError = err
			return
		}
		bodyBuf.Write(bodyData[0:readSize])
		if bodyBuf.Len() == bodyLen {
			break
		} else {
			bodyData = make([]byte, bodyLen-readSize)
		}
	}
	data = bodyBuf.Bytes()
	returnError = nil
	return
}
