package dynamic

import (
	"fmt"
	"time"
)

func FindClosestWord(entered string, list []string) string {
	startWork := time.Now()
	countOperation = 0

	matchList := make([]float64, 0, len(list))
	for _, l := range list {
		value := 0.0
		if len(entered) <= len(l) {
			value = float64(checkTwoWords(entered, l)) / float64(len(l))
		} else {
			value = float64(checkTwoWords(l, entered)) / float64(len(entered))
		}
		matchList = append(matchList, value)
	}

	max := matchList[0]
	indexMax := 0
	for i := 1; i < len(matchList); i++ {
		countOperation++
		if matchList[i] > max {
			max = matchList[i]
			indexMax = i
			countOperation++
		}
	}

	fmt.Println("DINAMIC PROGRAMMING (find the closest word)")
	fmt.Printf("amount checked words is %v\n", len(list))
	fmt.Printf("Result: %v\n", list[indexMax])
	fmt.Printf("number of operations = %d\n", countOperation)
	fmt.Printf("duration = %v ns (%v ms)\n\n",
		time.Since(startWork).Nanoseconds(),
		time.Since(startWork).Milliseconds())

	return list[indexMax]
}

func checkTwoWords(shortest, longest string) int {
	matrix := make([][]int, len(shortest))
	for i := range matrix {
		countOperation++
		matrix[i] = make([]int, len(longest))
	}

	for i := 0; i < len(shortest); i++ {
		equalInLine := false
		for j := 0; j < len(longest); j++ {
			countOperation++
			if shortest[i] == longest[j] && !equalInLine {
				equalInLine = true
				if i != 0 {
					if j != 0 {
						matrix[i][j] = matrix[i][j-1]
						if matrix[i-1][j] > matrix[i][j] {
							matrix[i][j] = matrix[i-1][j]
						}
						matrix[i][j]++
						countOperation++
						continue
					}
					matrix[i][j] = matrix[i-1][len(longest)-1] + 1
					countOperation++
					continue
				}
				if j != 0 {
					matrix[i][j] = matrix[i][j-1] + 1
					countOperation++
					continue
				}
				matrix[i][j] = 1
				countOperation++
				continue
			}

			if i != 0 {
				if j != 0 {
					matrix[i][j] = matrix[i][j-1]
					if matrix[i-1][j] > matrix[i][j] {
						matrix[i][j] = matrix[i-1][j]
					}
					countOperation++
					continue
				}
				matrix[i][j] = matrix[i-1][len(longest)-1]
				countOperation++
				continue
			}
			if j != 0 {
				matrix[i][j] = matrix[i][j-1]
				countOperation++
				continue
			}
			matrix[i][j] = 0
			countOperation++
			continue
		}
	}

	return matrix[len(shortest)-1][len(longest)-1]
}
