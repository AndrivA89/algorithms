package recursion

import (
	"fmt"
	"time"
)

func SumOfNumbers(array []int) int {
	startWork := time.Now()
	result := sumOfNumbers(array)

	fmt.Println("RECURSION")
	fmt.Printf("Sum %d elements of array is %d\n", len(array), result)
	fmt.Printf("duration = %v ms (%0.3v s)\n\n",
		time.Since(startWork).Milliseconds(),
		time.Since(startWork).Seconds())

	return result
}

func sumOfNumbers(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	return array[0] + sumOfNumbers(array[1:])
}
