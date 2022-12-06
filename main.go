package main

import (
	"github.com/AndrivA89/algorithms/dynamic"
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

	dynamic.FindBestOption(
		map[string]int{"Moscow": 2, "NewYork": 7, "London": 5, "Praga": 3},
		map[string]int{"Moscow": 1, "NewYork": 3, "London": 3, "Praga": 2}, 5)
	dynamic.FindBestOption(
		map[string]int{"Macbook": 8, "PS5": 9, "Samsung phone": 4, "Iphone": 3},
		map[string]int{"Macbook": 2999, "PS5": 1000, "Samsung phone": 1699, "Iphone": 1699}, 4999)
	dynamic.FindBestOption(
		map[string]int{"вода": 10, "книга": 3, "еда": 9, "куртка": 5, "камера": 6},
		map[string]int{"вода": 3, "книга": 1, "еда": 2, "куртка": 2, "камера": 1}, 6)

	dynamic.FindClosestWord("fesh", []string{"fish", "hash", "diss", "fresh", "ffresshh"})
	dynamic.FindClosestWord("fresh", []string{"fish", "hash", "diss", "f_false_re_word_sh", "fesh"})
	dynamic.FindClosestWord("ffresshh", []string{"fish", "hash", "diss", "fresh"})
}
