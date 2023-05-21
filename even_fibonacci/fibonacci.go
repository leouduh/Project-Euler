package main

import (
	"fmt"
)

func even_fibonacci(end_point int) ([]int, int) {
	x, y, sum := 1, 2, 0
	arr := []int{x}
	for y < end_point {
		arr = append(arr, y)
		if y%2 == 0 {
			sum += y
		}
		x, y = y, x+y
	}
	return arr, sum
}

func main() {
	series, sum := even_fibonacci(4_000_000)
	fmt.Println(series)
	fmt.Printf("The sum of all even fibonacci terms in the series above is %d", sum)
}
