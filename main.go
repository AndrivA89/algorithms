package main

import (
	"github.com/AndrivA89/algorithms/recursion"
	"github.com/AndrivA89/algorithms/search"
	"github.com/AndrivA89/algorithms/sorting"
	"github.com/AndrivA89/algorithms/tools"
)

func main() {
	sorting.SelectSort(tools.InitArray(1000))
	sorting.SelectSort(tools.InitArrayWithUniqElements(1000))
	sorting.SelectSort(tools.InitArray(50000))
	sorting.SelectSort(tools.InitArrayWithUniqElements(50000))

	sorting.QuickSort(tools.InitArray(1000))
	sorting.QuickSort(tools.InitArrayWithUniqElements(1000))
	sorting.QuickSort(tools.InitArray(50000))
	sorting.QuickSort(tools.InitArrayWithUniqElements(50000))
	sorting.QuickSort(tools.InitArray(1000000))
	sorting.QuickSort(tools.InitArrayWithUniqElements(1000000))

	recursion.Factorial(3)
	recursion.Factorial(50)

	recursion.SumOfNumbers(tools.InitArray(1000000))

	uniqArray := tools.InitArrayWithUniqElements(500000)
	search.BinarySearch(uniqArray, uniqArray[len(uniqArray)-1])
	uniqArray = tools.InitArrayWithUniqElements(500000)
	search.BinarySearch(uniqArray, tools.RandElementOfArray(uniqArray))

	uniqArray = tools.InitArrayWithUniqElements(500000)
	search.SearchByTree(uniqArray, uniqArray[len(uniqArray)-1])
	uniqArray = tools.InitArrayWithUniqElements(500000)
	search.SearchByTree(uniqArray, tools.RandElementOfArray(uniqArray))
}
