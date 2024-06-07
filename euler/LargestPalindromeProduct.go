package main

import (
	"fmt"
)

func revNumber(number int) int {
	reversedNumber := 0

	for number != 0 {
		reversedNumber *= 10
		lastDigit := number % 10
		reversedNumber += lastDigit
		number = number / 10
	}

	return reversedNumber
}

func main() {
	startingValue := 100

	biggestPalindrome := 0

	for i := 0; i < 899; i++ {
		tempValue := startingValue
		for j := 0; j < 899; j++ {
			value := (tempValue + j) * (startingValue + i)
			if value > biggestPalindrome {
				if value == revNumber(value) {
					biggestPalindrome = value
				}

			}

		}

	}
	fmt.Println(biggestPalindrome)

}
