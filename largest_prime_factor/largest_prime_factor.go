package main

import (
	"fmt"
	"math"
)

func checkPrime(num int) bool {
	//algorithm checks if number is a prime number given 6k +/- 1 to generate
	//prime numbers, 1 would be considered a prime number 2 and 3 are
	//prime numbers
	if num <= 3 && num != 1 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	n_sqrt := int(math.Sqrt(float64(num)))
	//5 and 7 would be true check their square root
	//for loop wouldn`t run so function ends up returning true
	for i := 5; i < n_sqrt+1; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func generatePrime(k int) (int, int) {
	primeNumber_1 := (6 * k) - 1
	primeNumber_2 := (6 * k) + 1
	return primeNumber_1, primeNumber_2
}

func largestPrimeFactor(number int) int {
	k := 0
	largestPn := 1
	largestPf := 1
	n_sqrt := int(math.Sqrt(float64(number)))
	for largestPn <= n_sqrt {
		n1, n2 := generatePrime(k)
		if checkPrime(n1) && n1 != -1 {
			largestPn = n1
			if number%largestPn == 0 {
				largestPf = largestPn
			}
		}
		if checkPrime(n2) {
			largestPn = n2
			if number%largestPn == 0 {
				largestPf = largestPn
			}
		}
		fmt.Println(largestPn)
		k++
	}
	return largestPf
}

func main() {
	number := 600851475143
	ans := largestPrimeFactor(number)
	fmt.Printf("The larglargestPn factor of %d is %d", number, ans)

}
