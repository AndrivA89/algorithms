package sorting

import (
	"fmt"
	"time"
)

func SelectSort(array []int) {
	startWork := time.Now()
	fmt.Printf("SELECT SORT - array size: %v\n", len(array))

	countOperation := 0
	for i := range array {
		min := i
		for j := i + 1; j < len(array); j++ {
			countOperation++
			if array[j] < array[min] {
				min = j
			}
		}

		if min != i {
			array[i], array[min] = array[min], array[i]
			countOperation++
		}
	}

	fmt.Printf("number of operations = %d\n", countOperation)
	fmt.Printf("duration = %v ms (%.3v s)\n\n",
		time.Since(startWork).Milliseconds(),
		time.Since(startWork).Seconds())
}
