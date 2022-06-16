package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	PhonePattern  = "(?:0|86|\\+86)?1[3-9]\\d{9}"
	EmailPattern  = "\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*"
	WeixinPattern = "^[a-zA-Z][a-zA-Z\\d_-]{5,19}$"
)

func IsValidMobile(phone string) bool {
	matched, err := regexp.Match(PhonePattern, []byte(phone))
	return err == nil && matched
}

func IsValidEmail(email string) bool {
	if strings.HasSuffix(email, ".") {
		return false
	}
	reg := regexp.MustCompile(EmailPattern)
	return reg.MatchString(email)
}

//
//func IsValidMobile(phone string) bool {
//	matched, err := regexp.Match(PhonePattern, []byte(phone))
//	return err == nil && matched
//}
//
//func IsValidEmail(email string) bool {
//	if strings.HasSuffix(email, ".") {
//		return false
//	}
//	reg := regexp.MustCompile(EmailPattern)
//	return reg.MatchString(email)
//}

func IsValidWeiXinId(weiXinId string) bool {
	matched, err := regexp.Match(WeixinPattern, []byte(weiXinId))
	//reg := regexp.MustCompile(WeixinPattern)
	//return reg.MatchString(weiXinId)
	return err == nil && matched
}

func main() {
	res := IsValidWeiXinId("")
	fmt.Println(res)
}
