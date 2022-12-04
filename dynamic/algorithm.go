package dynamic

import (
	"fmt"
	"time"
)

type item struct {
	name     string
	value    int // Значимость (важность)
	capacity int // Стоимость (длительность)
}

type result struct {
	value int
	items []string
}

var countOperation int

func setItems(itemsValue, itemsCapacity map[string]int) []item {
	result := make([]item, 0, len(itemsValue))
	for name, value := range itemsValue {
		countOperation++
		result = append(result, item{
			name:     name,
			value:    value,
			capacity: itemsCapacity[name],
		})
	}
	return result
}

// FindBestOption: itemsValue - значимость (важность),
// itemsCapacity - стоимость (длительность),
// capacity - доступный объем стоимости (длительности)
func FindBestOption(itemsValue, itemsCapacity map[string]int, capacity int) []string {
	startWork := time.Now()
	countOperation = 0

	items := setItems(itemsValue, itemsCapacity)

	// Первый столбец с нулевым индексом, поэтому (+1)
	amountStepsCapacity := capacity + 1

	// Создание таблицы весов
	matrix := make([][]result, len(items))
	for i := range matrix {
		countOperation++
		matrix[i] = make([]result, amountStepsCapacity)
	}

	// Заполняем первую строчку значением первого элемента
	for indexСolumn := 0; indexСolumn < amountStepsCapacity; indexСolumn++ {
		countOperation++
		if items[0].capacity <= indexСolumn {
			matrix[0][indexСolumn].value = items[0].value
			matrix[0][indexСolumn].items = []string{items[0].name}
		} else {
			matrix[0][indexСolumn].value = 0
			matrix[0][indexСolumn].items = []string{}
		}
	}

	for indexLine, i := range items {
		if indexLine == 0 {
			continue
		}
		for indexСolumn := 0; indexСolumn < amountStepsCapacity; indexСolumn++ {
			countOperation++
			// По умолчанию считаем максимумом значение на прошлой строке и таком же столбце
			max := matrix[indexLine-1][indexСolumn]
			matrix[indexLine][indexСolumn] = max
			if indexСolumn >= i.capacity {
				// Определяем новый возможный максимум, добавляя к текущему значению
				// значение, подходяще во весу из прошлой строки
				newMax := i.value + matrix[indexLine-1][indexСolumn-i.capacity].value
				if newMax > max.value {
					// Выбираем большее значение из двух максимумов
					matrix[indexLine][indexСolumn].value = newMax
					newItems := matrix[indexLine-1][indexСolumn-i.capacity].items
					if len(newItems) == 0 {
						newItems = []string{i.name}
					}

					// Не заполняем элемент повторно
					ok := true
					for _, n := range newItems {
						if i.name == n {
							ok = false
							break
						}
					}
					if ok {
						newItems = append(newItems, i.name)
					}

					matrix[indexLine][indexСolumn].items = newItems
				}
			}
		}
	}

	fmt.Println("DINAMIC PROGRAMMING (find best option)")
	fmt.Printf("amount options is %v\n", len(items))
	fmt.Printf("Result: %v\n", matrix[len(items)-1][amountStepsCapacity-1].items)
	fmt.Printf("number of operations = %d\n", countOperation)
	fmt.Printf("duration = %v ns (%v ms)\n\n",
		time.Since(startWork).Nanoseconds(),
		time.Since(startWork).Milliseconds())
	return matrix[len(items)-1][amountStepsCapacity-1].items
}
