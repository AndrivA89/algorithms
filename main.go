package main

import (
	"github.com/AndrivA89/algorithms/sorting"
)

var arr []int

func main() {
	arr = initArray(1000)
	sorting.SelectSort(arr)

	arr = initArray(10000)
	sorting.SelectSort(arr)

	arr = initArray(100000)
	sorting.SelectSort(arr)
}
