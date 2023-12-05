package day2

import (
	"strconv"
	"strings"
)

func PartOne() int64 {
	maxRedCubes := 12
	maxGreenCubes := 13
	maxBlueCubes := 14

	var possibleGamesCount int64
	for _, input := range inputData {
		in := strings.Split(input, ": ")
		subsets := strings.Split(in[1], "; ")

		isPossible := true
		for _, subset := range subsets {
			counts := strings.Split(subset, ", ")
			for _, count := range counts {
				var blueCount int64
				if strings.Contains(count, "blue") {
					n := count[:(len(count) - 5)]
					blueCount, _ = strconv.ParseInt(n, 10, 32)
					isPossible = blueCount <= int64(maxBlueCubes)
					if !isPossible {
						break
					}
				}

				var redCount int64
				if strings.Contains(count, "red") {
					n := count[:(len(count) - 4)]
					redCount, _ = strconv.ParseInt(n, 10, 32)
					isPossible = redCount <= int64(maxRedCubes)
					if !isPossible {
						break
					}
				}

				var greenCount int64
				if strings.Contains(count, "green") {
					n := count[:(len(count) - 6)]
					greenCount, _ = strconv.ParseInt(n, 10, 32)
					isPossible = greenCount <= int64(maxGreenCubes)
					if !isPossible {
						break
					}
				}
			}
			if !isPossible {
				break
			}
		}

		if isPossible {
			n := strings.Split(in[0], "Game ")
			gameCount, _ := strconv.ParseInt(n[1], 10, 32)
			possibleGamesCount += gameCount
		}
	}

	return possibleGamesCount
}

func PartTwo() int64 {
	var sumCubes int64
	for _, input := range inputData {
		// game
		in := strings.Split(input, ": ")
		subsets := strings.Split(in[1], "; ")

		var minBlue int64
		var minGreen int64
		var minRed int64
		for _, subset := range subsets {
			// rounds
			counts := strings.Split(subset, ", ")
			for _, count := range counts {
				// cubes
				var blueCount int64
				if strings.Contains(count, "blue") {
					n := count[:(len(count) - 5)]
					blueCount, _ = strconv.ParseInt(n, 10, 32)
				}
				if minBlue < blueCount {
					minBlue = blueCount
				}

				var redCount int64
				if strings.Contains(count, "red") {
					n := count[:(len(count) - 4)]
					redCount, _ = strconv.ParseInt(n, 10, 32)
				}
				if minRed < redCount {
					minRed = redCount
				}

				var greenCount int64
				if strings.Contains(count, "green") {
					n := count[:(len(count) - 6)]
					greenCount, _ = strconv.ParseInt(n, 10, 32)
				}
				if minGreen < greenCount {
					minGreen = greenCount
				}
			}
		}
		sumCubes += minBlue * minGreen * minRed
	}
	return sumCubes
}
