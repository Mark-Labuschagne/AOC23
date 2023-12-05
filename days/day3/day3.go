package day3

import "strconv"

func PartOne() int64 {
	var sum int64
	for i, line := range inputData {
		for j := 0; j < len(line); j++ {
			r := rune(line[j])

			if !isNumber(r) {
				continue
			}

			if i > 0 {
				// check above
				if isSymbol(rune(inputData[i-1][j])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
			if j > 0 {
				// check left
				if isSymbol(rune(inputData[i][j-1])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
			if i < len(inputData)-1 {
				// check below
				if isSymbol(rune(inputData[i+1][j])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
			if j < len(line)-1 {
				// check right
				if isSymbol(rune(inputData[i][j+1])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}

			if i > 0 && j > 0 {
				// check above-left
				if isSymbol(rune(inputData[i-1][j-1])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
			if i > 0 && j < len(line)-1 {
				// check above-right
				if isSymbol(rune(inputData[i-1][j+1])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
			if i < len(inputData)-1 && j > 0 {
				// check below-left
				if isSymbol(rune(inputData[i+1][j-1])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
			if i < len(inputData)-1 && j < len(line)-1 {
				// check below-left
				if isSymbol(rune(inputData[i+1][j+1])) {
					number, lastIndex := getWholeNumber(inputData[i], j)
					sum += number
					j = lastIndex + 1
					continue
				}
			}
		}
	}

	return sum
}

func isNumber(r rune) bool {
	return r >= []rune("0")[0] && r <= []rune("9")[0]
}

func isSymbol(r rune) bool {
	return !isNumber(r) && r != []rune(".")[0]
}

// getWholeNumber returns the number, and the index of the last number in wholeNumber
func getWholeNumber(s string, startIndex int) (int64, int) {
	var wholeNumber int64
	if startIndex == 0 {
		var number string
		for i, ss := range s {
			if isNumber(ss) {
				number += string(ss)
				continue
			}

			wholeNumber, _ = strconv.ParseInt(number, 10, 64)
			return wholeNumber, i - 1
		}
	}

	for {
		// walk backwards until we've found the start of the number
		if startIndex-1 == -1 {
			break
		} else if !isNumber(rune(s[startIndex-1])) {
			break
		}

		startIndex--
	}

	var number string
	for i := startIndex; i < len(s); i++ {
		if isNumber(rune(s[i])) {
			number += string(s[i])
			continue
		}

		wholeNumber, _ = strconv.ParseInt(number, 10, 64)
		return wholeNumber, i - 1
	}

	wholeNumber, _ = strconv.ParseInt(number, 10, 64)
	return wholeNumber, len(s) - 1
}

func PartTwo() int64 {
	var mulTotal int64
	for i, line := range inputData {
		var mul int64 = 1
		checkedNumbersMap := make(map[int64]bool)
		for j := 0; j < len(line); j++ {
			r := rune(line[j])

			if !isGear(r) {
				continue
			}

			var adjacentCount int
			if i > 0 {
				// check above
				if isNumber(rune(inputData[i-1][j])) {
					number, _ := getWholeNumber(inputData[i-1], j)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}
			if j > 0 {
				// check left
				if isNumber(rune(inputData[i][j-1])) {
					number, _ := getWholeNumber(inputData[i], j-1)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}
			if i < len(inputData)-1 {
				// check below
				if isNumber(rune(inputData[i+1][j])) {
					number, _ := getWholeNumber(inputData[i+1], j)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}
			if j < len(line)-1 {
				// check right
				if isNumber(rune(inputData[i][j+1])) {
					number, _ := getWholeNumber(inputData[i], j+1)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}

			if i > 0 && j > 0 {
				// check above-left
				if isNumber(rune(inputData[i-1][j-1])) {
					number, _ := getWholeNumber(inputData[i-1], j-1)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}
			if i > 0 && j < len(line)-1 {
				// check above-right
				if isNumber(rune(inputData[i-1][j+1])) {
					number, _ := getWholeNumber(inputData[i-1], j+1)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}
			if i < len(inputData)-1 && j > 0 {
				// check below-left
				if isNumber(rune(inputData[i+1][j-1])) {
					number, _ := getWholeNumber(inputData[i+1], j-1)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
						checkedNumbersMap[number] = true
					}
				}
			}
			if i < len(inputData)-1 && j < len(line)-1 {
				// check below-left
				if isNumber(rune(inputData[i+1][j+1])) {
					number, _ := getWholeNumber(inputData[i+1], j+1)
					if !checkedNumbersMap[number] {
						mul *= number
						adjacentCount++
					}
				}
			}

			if adjacentCount == 2 {
				mulTotal += mul
			}
			mul = 1
		}
	}

	return mulTotal
}

func isGear(r rune) bool {
	return r == []rune("*")[0]
}
