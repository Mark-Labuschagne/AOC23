package day1

import (
	"strconv"
)

func Calibrate() int {
	lowerLimit := []rune("0")[0]
	upperLimit := []rune("9")[0]
	var total int
	for _, input := range inputData {
		// get first number
		var firstNumber string
		for _, r := range input {
			if r >= lowerLimit && r <= upperLimit {
				firstNumber = string(r)
				break
			}
		}

		var lastNumber string
		for i := len(input) - 1; i > -1; i-- {
			r := rune(input[i])
			if r >= lowerLimit && r <= upperLimit {
				lastNumber = string(r)
				break
			}
		}
		thisTotal, err := strconv.ParseInt(firstNumber+lastNumber, 10, 64)
		if err != nil {
			panic(err)
		}
		total += int(thisTotal)
	}

	return total
}

func CalibrateWithWords() int {
	lowerLimit := []rune("0")[0]
	upperLimit := []rune("9")[0]
	var total int
	for _, input := range inputData {
		// get first number
		var firstNumber string
		for i, r := range input {
			lenS := len(input)

			if r >= lowerLimit && r <= upperLimit {
				firstNumber = string(r)
				break
			}

			if r == 'o' && i+2 <= lenS-1 && input[i+1] == 'n' && input[i+2] == 'e' {
				firstNumber = "1"
				break
			}
			if r == 't' && i+2 <= lenS-1 && input[i+1] == 'w' && input[i+2] == 'o' {
				firstNumber = "2"
				break
			}
			if r == 't' && i+4 <= lenS-1 && input[i+1] == 'h' && input[i+2] == 'r' && input[i+3] == 'e' && input[i+4] == 'e' {
				firstNumber = "3"
				break
			}
			if r == 'f' && i+3 <= lenS-1 && input[i+1] == 'o' && input[i+2] == 'u' && input[i+3] == 'r' {
				firstNumber = "4"
				break
			}
			if r == 'f' && i+3 <= lenS-1 && input[i+1] == 'i' && input[i+2] == 'v' && input[i+3] == 'e' {
				firstNumber = "5"
				break
			}
			if r == 's' && i+2 <= lenS-1 && input[i+1] == 'i' && input[i+2] == 'x' {
				firstNumber = "6"
				break
			}
			if r == 's' && i+4 <= lenS-1 && input[i+1] == 'e' && input[i+2] == 'v' && input[i+3] == 'e' && input[i+4] == 'n' {
				firstNumber = "7"
				break
			}
			if r == 'e' && i+4 <= lenS-1 && input[i+1] == 'i' && input[i+2] == 'g' && input[i+3] == 'h' && input[i+4] == 't' {
				firstNumber = "8"
				break
			}
			if r == 'n' && i+3 <= lenS-1 && input[i+1] == 'i' && input[i+2] == 'n' && input[i+3] == 'e' {
				firstNumber = "9"
				break
			}
		}

		var lastNumber string
		for i := len(input) - 1; i > -1; i-- {
			r := rune(input[i])
			if r >= lowerLimit && r <= upperLimit {
				lastNumber = string(r)
				break
			}

			if r == 'e' && i-2 >= 0 && input[i-1] == 'n' && input[i-2] == 'o' {
				lastNumber = "1"
				break
			}
			if r == 'o' && i-2 >= 0 && input[i-1] == 'w' && input[i-2] == 't' {
				lastNumber = "2"
				break
			}
			if r == 'e' && i-4 >= 0 && input[i-1] == 'e' && input[i-2] == 'r' && input[i-3] == 'h' && input[i-4] == 't' {
				lastNumber = "3"
				break
			}
			if r == 'r' && i-3 >= 0 && input[i-1] == 'u' && input[i-2] == 'o' && input[i-3] == 'f' {
				lastNumber = "4"
				break
			}
			if r == 'e' && i-3 >= 0 && input[i-1] == 'v' && input[i-2] == 'i' && input[i-3] == 'f' {
				lastNumber = "5"
				break
			}
			if r == 'x' && i-2 >= 0 && input[i-1] == 'i' && input[i-2] == 's' {
				lastNumber = "6"
				break
			}
			if r == 'n' && i-4 >= 0 && input[i-1] == 'e' && input[i-2] == 'v' && input[i-3] == 'e' && input[i-4] == 's' {
				lastNumber = "7"
				break
			}
			if r == 't' && i-4 >= 0 && input[i-1] == 'h' && input[i-2] == 'g' && input[i-3] == 'i' && input[i-4] == 'e' {
				lastNumber = "8"
				break
			}
			if r == 'e' && i-3 >= 0 && input[i-1] == 'n' && input[i-2] == 'i' && input[i-3] == 'n' {
				lastNumber = "9"
				break
			}

		}
		thisTotal, err := strconv.ParseInt(firstNumber+lastNumber, 10, 64)
		if err != nil {
			panic(err)
		}
		total += int(thisTotal)
	}

	return total
}
