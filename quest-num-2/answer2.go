package main

var numString string = "23242526272830"
var num1, num2 int
var max int = 6
var arr []int

func getValue(str string, i int, m int) (value int) {

	if i+m > len(str) {
		return -1
	}

	value = 0
	for j := 0; j < m; j++ {
		var c = str[i+j] - '0'
		if c < 0 || c > 9 {
			return -1
		}

		value = value * 10
	}
	return value
}

func main() {

	getValue(numString, 1, 12)
}

// not completed
