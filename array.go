package main

import (
	"math/rand"
	"time"
)

func initArray(size int) []int {
	array := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		minus := -1
		if r.Intn(2) == 1 {
			minus = 1
		}
		array[i] = r.Intn(100) * minus
	}

	return array
}
