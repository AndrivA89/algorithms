package sorting

import (
	"fmt"
	"time"
)

var countOperation int

func QuickSort(array []int) {
	startWork := time.Now()
	countOperation = 0

	quickSort(array, 0, len(array)-1)

	fmt.Printf("QUICK SORT - array size: %v\n", len(array))
	fmt.Printf("number of operations = %d\n", countOperation)
	fmt.Printf("duration = %v ms (%.3v s)\n\n",
		time.Since(startWork).Milliseconds(),
		time.Since(startWork).Seconds())
}

func quickSort(data []int, start, end int) {
	if start < end {
		base := data[start]
		left := start
		right := end

		for left < right {
			for left < right && data[right] >= base {
				right--
				countOperation++
			}
			if left < right {
				data[left] = data[right]
				left++
				countOperation++
			}

			for left < right && data[left] <= base {
				left++
				countOperation++
			}
			if left < right {
				data[right] = data[left]
				right--
				countOperation++
			}
		}

		data[left] = base
		countOperation++

		quickSort(data, start, left-1)
		quickSort(data, left+1, end)
	}
}
