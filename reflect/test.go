package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"

	"net/http"
	"reflect"
)

type BaseResp struct {
	StatusMessage string            `thrift:"StatusMessage,1" json:"StatusMessage"`
	StatusCode    int32             `thrift:"StatusCode,2" json:"StatusCode"`
	Extra         map[string]string `thrift:"Extra,3" json:"Extra,omitempty"`
}

type Response struct {
	Data     *bool     `thrift:"data,1" json:"data,omitempty"`
	Code     *int64    `thrift:"code,253" json:"code,omitempty"`
	Message  *string   `thrift:"message,254" json:"message,omitempty"`
	BaseResp *BaseResp `thrift:"BaseResp,255" json:"BaseResp,omitempty"`
}

func ConvertErrorCode(base *BaseResp, rsp interface{}) {
	if base == nil {
		reflect.ValueOf(rsp).Elem().FieldByName("Code").SetInt(0)
		reflect.ValueOf(rsp).Elem().FieldByName("Message").SetString("成功")
	} else if base.StatusCode != 0 {
		reflect.ValueOf(rsp).Elem().FieldByName("Code").SetInt(int64(http.StatusBadRequest))
		reflect.ValueOf(rsp).Elem().FieldByName("Message").SetString(base.StatusMessage)
	} else {
		reflect.ValueOf(rsp).Elem().FieldByName("Code").SetInt(int64(base.StatusCode))
		reflect.ValueOf(rsp).Elem().FieldByName("Message").SetString("成功")
	}
}

func main() {
	var r interface{}
	r = &Response{
		Code: aws.Int64(0),
		BaseResp: &BaseResp{
			StatusCode: 0,
		},
	}
	t := reflect.TypeOf(r)
	fmt.Println(t.Name(), t.Kind())
	v := reflect.ValueOf(r)
	//fmt.Println(t.NumField())
	elem := v.Elem()
	//for i := 0; i < t.NumField(); i++ {
	//	// Field 代表对象的字段名
	//	key := t.Field(i)
	//	value := v.Field(i).Interface()
	//	// 字段
	//	if key.Anonymous {
	//		fmt.Printf("匿名字段 第 %d 个字段，字段名 %s, 字段类型 %v, 字段的值 %v\n", i+1, key.Name, key.Type, value)
	//	} else {
	//		fmt.Printf("命名字段 第 %d 个字段，字段名 %s, 字段类型 %v, 字段的值 %v\n", i+1, key.Name, key.Type, value)
	//	}
	//}
	if elem.FieldByName("Code").Kind().String() == "ptr" {
		elem.FieldByName("Code").Set(reflect.ValueOf(aws.Int64(9)))
	} else {
		elem.FieldByName("Code").SetInt(0)
	}
	baseResp := elem.FieldByName("BaseResp").Elem()
	fmt.Println(baseResp.CanSet())
	//baseResp.FieldByName("StatusCode").SetInt(1)
	fmt.Println(baseResp.IsZero())
	//fmt.Println(baseResp.FieldByName("StatusCode").Int())
}
