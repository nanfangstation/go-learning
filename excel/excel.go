package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	file, err := excelize.OpenFile("/Users/xx模板.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//// 获取工作表中指定单元格的值
	//cell, err := file.GetColOutlineLevel("Sheet1", "1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows[1])
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}
	if validHeader := ContainsString(rows[1], "ID") &&
		ContainsString(rows[1], "抖音POIID") &&
		ContainsString(rows[1], "经度") &&
		ContainsString(rows[1], "纬度") &&
		ContainsString(rows[1], "门店名称") &&
		ContainsString(rows[1], "省") &&
		ContainsString(rows[1], "市") &&
		ContainsString(rows[1], "区") &&
		ContainsString(rows[1], "地址") &&
		ContainsString(rows[1], "行业类别") &&
		ContainsString(rows[1], "门店电话") &&
		ContainsString(rows[1], "填报人电话") &&
		ContainsString(rows[1], "营业状态"); !validHeader {
		fmt.Println("Excel模板格式错误，请下载最新模板后重新提交")
	}
	fmt.Println(ContainsString(rows[1], "ID"))
	fmt.Println(rows[2][1])
	fmt.Println("正常结束")
}

func ContainsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}
