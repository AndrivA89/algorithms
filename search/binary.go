package search

import (
	"fmt"
	"time"

	"github.com/AndrivA89/algorithms/sorting"
)

func BinarySearch(array []int, target int) {
	fmt.Printf("BINARY SEARCH for %d elements of array\n", len(array))
	countOperation = 0

	startSort := time.Now()
	sorting.QuickSortWithoutLogs(array)

	startSearch := time.Now()

	left, right := 0, len(array)-1

	for left <= right {
		center := (left + right) / 2
		countOperation++
		if array[center] > target {
			right = center - 1
			countOperation++
		} else if array[center] < target {
			left = center + 1
			countOperation++
		} else {
			fmt.Printf("target element's '%d' index is %d\n", target, center)
			fmt.Printf("duration sort: %v ms (%v ns)\n",
				time.Since(startSort).Milliseconds(), time.Since(startSort).Nanoseconds())
			fmt.Printf("duration search: %v ms (%v ns)\n",
				time.Since(startSearch).Milliseconds(), time.Since(startSearch).Nanoseconds())
				fmt.Printf("number of operations = %d\n\n", countOperation)
			return
		}
	}

	fmt.Printf("BINARY SEARCH element %d isn't found\n", target)
	fmt.Printf("duration sort: %v ms (%v ns)\n",
		time.Since(startSort).Milliseconds(), time.Since(startSort).Nanoseconds())
	fmt.Printf("duration search: %v ms (%v ns)\n",
		time.Since(startSearch).Milliseconds(), time.Since(startSearch).Nanoseconds())
		fmt.Printf("number of operations = %d\n\n", countOperation)
}
