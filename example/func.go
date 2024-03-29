package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	return !(integer%2 == 0)
}

func isEven(integer int) bool {
	return integer%2 == 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice=", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are:", odd)
}
