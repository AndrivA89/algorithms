package main

import (
	"github.com/AndrivA89/algorithms/recursion"
	"github.com/AndrivA89/algorithms/sorting"
	"github.com/AndrivA89/algorithms/tools"
)

func main() {
	sorting.SelectSort(tools.InitArray(1000))
	sorting.SelectSort(tools.InitArray(50000))

	sorting.QuickSort(tools.InitArray(1000))
	sorting.QuickSort(tools.InitArray(50000))
	sorting.QuickSort(tools.InitArray(1000000))

	recursion.Factorial(3)
	recursion.Factorial(50)

	recursion.SumOfNumbers(tools.InitArray(1000000))
}
