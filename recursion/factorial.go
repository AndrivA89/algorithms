package recursion

import (
	"fmt"
	"time"
)

func factorial(number uint64) uint64 {
	if number == 1 {
		return number
	}

	return number * factorial(number-1)
}

func Factorial(number uint64) uint64 {
	now := time.Now()
	result := factorial(number)

	fmt.Printf("FACTORIAL %d is %d\n", number, result)
	fmt.Printf("duration %v ns\n\n", time.Since(now).Nanoseconds())

	return result
}
