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

func main() {
	all_primes := []int{2, 3}
	k := 1
	for len(all_primes) <= 10001 {
		n1, n2 := generatePrime(k)
		if checkPrime(n1) {
			all_primes = append(all_primes, n1)
		}
		if checkPrime(n2) {
			all_primes = append(all_primes, n2)
		}
		k += 1
	}
	fmt.Printf("The 100001st prime number is %d", all_primes[10000])

}
