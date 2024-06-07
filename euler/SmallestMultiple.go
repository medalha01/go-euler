package main

import (
	"fmt"
)

func isDivisibleToN(num int, n int) bool {
	for i := 1; i < n; i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	startingPoint := 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19

	for true {
		if isDivisibleToN(startingPoint, 20) {
			fmt.Println(startingPoint)
			return
		}
		startingPoint += 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19

	}

}
