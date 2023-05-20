package main

import (
	"fmt"
)

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func sum_squares_n(n int) int {
	return (n * (n + 1) * ((2 * n) + 1)) / 6
}

func create_arr(n int) []int {
	arr := []int{}
	i := 1
	for i <= n {
		arr = append(arr, i)
		i++
	}
	return arr
}
func main() {
	n := 100
	series := create_arr(n)
	sumOfSeries := sum(series)
	fmt.Printf("The sum of the sum square difference of the first %d natural numbers is %d", n, (sumOfSeries*sumOfSeries)-sum_squares_n(len(series)))
}
