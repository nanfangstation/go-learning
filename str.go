package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var rules *[]Rule
	ruleStr := "[{\"protocol\":\"HTTP\",\"port\":\"8080\",\"domain\":[\"httpbin.org\"]}]"
	if err := json.Unmarshal([]byte(ruleStr), &rules); err != nil {
		fmt.Println(err)
	}
	fmt.Println(rules)
}

type MsgStruct struct {
	Type      int32   `json:"type"` // 1-集群 2-规则
	ClusterId string  `json:"cluster_id"`
	AccountId int64   `json:"account_id"`
	AK        string  `json:"ak"`
	SK        string  `json:"sk"`
	Rules     *[]Rule `json:"rules"`
}

type Rule struct {
	Protocol string   `json:"protocol"`
	Domain   []string `json:"domain"`
	Port     string   `json:"port"`
}
