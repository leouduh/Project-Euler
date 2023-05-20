package main

import (
	"fmt"
)

func multiplyAll(arr []int) int {
	total := 1
	for _, val := range arr {
		total *= val
	}
	return total
}

func main() {
	var factors_slice = []int{}
	num := 2
	endpoint := 20
	for num <= endpoint {
		if len(factors_slice) == 0 {
			factors_slice = append(factors_slice, num)
			num += 1
			continue
		}
		number := num
		for index, factor := range factors_slice {
			if number%factor == 0 {
				number /= factor
			}
			if number == 1 {
				continue
			}
			if number != 1 && index == len(factors_slice)-1 {
				factors_slice = append(factors_slice, number)
			}
		}
		num += 1
	}
	fmt.Println(multiplyAll(factors_slice))
}
