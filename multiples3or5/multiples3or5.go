package main

import (
	"fmt"
)

func multiplesOf3or5(endpoint int) int {
	sum := 0
	for i := 3; i < endpoint; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 && i >= 5 {
			sum += i
		}
	}
	return sum
}

func main() {
	num := 1000
	fmt.Printf("the sum of all multiples of 3 and 5 between zero and %d is %d", num, multiplesOf3or5(num))
}
