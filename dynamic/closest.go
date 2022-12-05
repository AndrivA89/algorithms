package dynamic

import (
	"fmt"
	"time"
)

func FindClosestWord(entered string, list []string) string {
	startWork := time.Now()
	countOperation = 0

	matchList := make([]int, 0, len(list))
	for _, l := range list {
		value := checkTwoWords(entered, l)
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

func checkTwoWords(entered, forCheck string) int {
	matrix := make([][]int, len(entered))
	for i := range matrix {
		countOperation++
		matrix[i] = make([]int, len(forCheck))
	}

	for i := 0; i < len(entered); i++ {
		for j := 0; j < len(forCheck); j++ {
			countOperation++
			if i > len(forCheck) {
				break
			}
			if entered[i] == forCheck[j] {
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
					matrix[i][j] = matrix[i-1][j] + 1
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
				matrix[i][j] = matrix[i-1][j]
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

	return matrix[len(entered)-1][len(forCheck)-1]
}
