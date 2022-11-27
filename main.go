package main

import (
	"fmt"

	"github.com/AndrivA89/algorithms/sorting"
)

func Sort() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, 8, -9}

	result := sorting.SelectSort(arr)
	fmt.Println(result)
}
