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
	startWork := time.Now()
	result := factorial(number)

	fmt.Println("RECURSION")
	fmt.Printf("Factorial %d is %d\n", number, result)
	fmt.Printf("duration %v ns\n\n", time.Since(startWork).Nanoseconds())

	return result
}
