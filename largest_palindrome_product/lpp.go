package main

import (
	"fmt"
	"strconv"
	"strings"
	// "math"
)

func break_arr(arr []int) ([]int, []int) {
	var x int = len(arr)
	midPoint := int(len(arr) / 2)
	first_arr := arr[0:midPoint]
	second_arr := arr[midPoint:x]
	return first_arr, second_arr
}

func checkPalindrome(num int) bool {
	var status bool
	num_string_arr := strings.Split(strconv.Itoa(num), "")
	x := len(num_string_arr)
	for index := 0; index < x; index++ {
		if num_string_arr[index] == num_string_arr[x-index-1] {
			status = true
		} else {
			status = false
			break
		}
	}
	return status
}

func create_arr(max_number int) []int {
	all_num := []int{}
	i := 1
	for i < max_number {
		all_num = append(all_num, i)
		i++
	}
	return all_num
}

func max(lst []int) int {
	maximum := lst[0]
	for i := 1; i < len(lst); i++ {
		if lst[i] > maximum {
			maximum = lst[i]
		}
	}
	return maximum
}

var products = []int{}
var prod int

func lpp(all_num []int) {
	if len(all_num) >= 2 {
		first_arr, second_arr := break_arr(all_num)
		for _, val1 := range first_arr {
			for _, val2 := range second_arr {
				prod = val1 * val2
				if checkPalindrome(prod) {
					products = append(products, prod)
				}
			}
		}
		lpp(first_arr)
		lpp(second_arr)
	}

}
func main() {
	max_number := 1000
	lst := create_arr(max_number)
	lpp(lst)
	fmt.Println("done")
	fmt.Printf("The largest palindrome product of two numbers between 0 and %d included is %d \n\n\n", max_number-1, max(products))
}
