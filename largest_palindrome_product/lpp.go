package main

import (
	"fmt"
	"strconv"
)

// func lpp(){
// 	decimals := 0
// 	tens := 0
// 	hundreds := 0

// }
func main() {
	number := 10
	num_str := strconv.Itoa(number)
	fmt.Printf("%T ---> %T", number, num_str)

}
