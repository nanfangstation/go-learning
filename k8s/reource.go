package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"k8s.io/apimachinery/pkg/api/resource"
	"math"
	"strconv"
)

func main() {
	println(convertCpu("200m"))
	println(convertMem("7908084Ki"))
}

func convertCpu(value string) *string {
	println("cpu值1:", value)
	cpu, _ := resource.ParseQuantity(value)
	println("cpu值2:", cpu.AsApproximateFloat64())
	println("cpu值3:", float64(cpu.ScaledValue(-3))/math.Pow10(3))
	return aws.String(strconv.FormatFloat(float64(cpu.ScaledValue(-3))/math.Pow10(3), 'f', 2, 64))
}

func convertMem(value string) *string {
	println("men值1:", value)
	mem, _ := resource.ParseQuantity(value)
	println("mem值2: ", mem.Value())
	println("mem值4: ", strconv.FormatFloat(float64(mem.Value())/math.Pow(2.0, 30.0), 'f', 2, 64))
	return aws.String(strconv.FormatFloat(float64((mem.Value())>>30), 'f', 2, 64))
}
