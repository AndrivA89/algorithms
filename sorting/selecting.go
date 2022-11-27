package sorting

func SelectSort(array []int) []int {
	var result []int

	for i := range array {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}

	return result
}
