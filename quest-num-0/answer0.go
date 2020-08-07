package main

import "fmt"

// set variable for get 2 number input
var number1 = 21
var number2 = 31

// set variable total palindrome
var totalPalindrome int

func main() {

	// set input value for number1
	fmt.Print("input number1: \n")
	fmt.Scan(&number1)

	// set input value for number2
	fmt.Print("input number2: \n")
	fmt.Scan(&number2)

	// check if first number is < than second number
	if number1 < number2 {
		// looping numbers based on user input
		for i := number1; i < number2; i++ {
			// set value for oriNumber and revNumber
			oriNumber := i
			revNumber := 0

			num := oriNumber

			for num > 0 {
				var remainder = num % 10
				revNumber = (revNumber * 10) + remainder
				num = num / 10
			}

			// the process of checking the palindrome number
			if oriNumber == revNumber {
				totalPalindrome++
			}
		}

		// showing total palindrome number between number1 and number2
		fmt.Println(totalPalindrome)
	} else {
		fmt.Printf("cannot process palindrome numbers between %d and %d", number1, number2)
		fmt.Print("\n")
	}
}
