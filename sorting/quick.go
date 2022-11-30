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

func QuickSortWithoutLogs(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(data []int, start, end int) {
	if start < end {
		base := data[start] // опорный элемент
		left := start
		right := end

		for left < right {
			// Выбираем первый элемент меньше базового, уменьшая правую границу
			for left < right && data[right] >= base {
				right--
				countOperation++
			}
			if left < right {
				// Закидываем элемент, меньший базового в левую часть массива в нашем случае на место опорного элемента
				data[left] = data[right]
				left++
				countOperation++
			}

			// Выбираем первый элемент больше базового, увеличивая левую границу
			for left < right && data[left] <= base {
				left++
				countOperation++
			}
			if left < right {
				// Закидываем элемент, больший базового в правую часть массива на место бывшего элемента меньшего опорного
				data[right] = data[left]
				right--
				countOperation++
			}
		}

		// Устанавливаем опорный элемент на место первого элемента больше опорного, найденного от левой границы
		data[left] = base
		countOperation++

		// Начинаем рекурсивно делить массив выбранным опорным элементом
		quickSort(data, start, left-1)
		quickSort(data, left+1, end)
	}
}
