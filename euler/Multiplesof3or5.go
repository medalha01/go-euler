package main

import (
	"fmt"
)

/*func main() {
	var queue [1000]int
	j := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			queue[j] = i
			j++
		}
	}
	sum := 0
	for _, value := range queue {
		sum += value
	}

	fmt.Println(sum)
}*/

// Function to calculate the sum of multiples of a number below a limit
func sumOfMultiples(factor, limit int) int {
	n := (limit - 1) / factor
	return factor * n * (n + 1) / 2
}

func main() {
	limit := 1000
	sum := sumOfMultiples(3, limit) + sumOfMultiples(5, limit) - sumOfMultiples(15, limit)
	fmt.Println(sum)
}
