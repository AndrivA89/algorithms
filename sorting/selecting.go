package sorting

import (
	"fmt"
	"time"
)

func SelectSort(array []int) {
	startWork := time.Now()
	fmt.Printf("SELECT SORT - array size: %v\n", len(array))

	countSwap := 0
	for i := range array {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		if min != i {
			array[i], array[min] = array[min], array[i]
			countSwap++
		}
	}

	fmt.Printf("number of operations = %d\n", countSwap)
	fmt.Printf("duration = %vms (%0.3vs)\n\n",
		time.Since(startWork).Milliseconds(),
		time.Since(startWork).Seconds())
}
