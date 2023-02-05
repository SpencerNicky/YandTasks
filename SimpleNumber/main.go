package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimes(15))
}

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(n-1))); i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	count := 1
	for i := 3; i < n; i += 2 {
		if primes[i] {
			count++
		}
	}
	return count
}
