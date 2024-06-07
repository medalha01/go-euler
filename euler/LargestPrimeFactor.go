package main

import (
	"fmt"
	"math"
)

func sieveOfEratosthenes(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
	}

	primes := []int{}
	for p := 2; p <= limit; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func contains(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
		if a > num {
			return false
		}
	}
	return false
}

func factorMultiples(num []int) int {
	mult_counter := 1
	for _, a := range num {
		mult_counter *= a
	}
	return mult_counter

}

func main() {
	originalNumber := 600851475143
	limit := int(math.Sqrt(float64(originalNumber))) + 1
	primes := sieveOfEratosthenes(limit)
	factor := []int{}
	remaining := originalNumber

	for _, prime := range primes {
		if prime*prime > remaining {
			break
		}
		for remaining%prime == 0 {
			factor = append(factor, prime)
			remaining /= prime
		}
	}

	if remaining > 1 {
		factor = append(factor, remaining)
	}

	fmt.Println("Prime factors:", factor)
}
