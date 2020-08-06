package main

import (
	"fmt"
	"testing"
)

// testing for main function answer0.go
func TestMain(t *testing.T) {

	if number1 == 1 && number2 == 10 {
		// unit test 1
		main()
		if totalPalindrome == 9 {
			fmt.Println("unit test 1 with result :", totalPalindrome)
		} else {
			t.Logf("Got %d expect %d", totalPalindrome, 9)
			t.Fail()
		}
	}

	if number1 == 99 && number2 == 100 {
		// unit test 2
		main()
		if totalPalindrome == 1 {
			fmt.Println("unit test 2 with result :", totalPalindrome)
		} else {
			t.Logf("Got %d expect %d", totalPalindrome, 1)
			t.Fail()
		}
	}

	if number1 == 21 && number2 == 31 {
		// unit test 3
		main()
		if totalPalindrome == 1 {
			fmt.Println("unit test 3 with result :", totalPalindrome)
		} else {
			t.Logf("Got %d expect %d", totalPalindrome, 1)
			t.Fail()
		}
	}

	if number1 > number2 {
		// other conditional
		t.Logf("first number is : %d must be < than %d", number1, number2)
		t.Fatal()
	} else if number1 == number2 {
		t.Log("first number cannot be == ", number2)
		t.Fail()
	} else {
		// testing successfully!
	}
}
