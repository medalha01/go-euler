package main

import (
	"fmt"
)

func getFibonnaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return getFibonnaci(n-1) + getFibonnaci(n-2)
}

func main() {
	numbers := [2]int{0, 1}

	value := 0
	sum := 0

	for value < 4000000 {
		value = numbers[0] + numbers[1]
		numbers[0] = numbers[1]
		numbers[1] = value
		if value%2 == 0 {
			sum += value
		}
	}

	fmt.Println(sum)
}
