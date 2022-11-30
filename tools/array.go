package tools

import (
	"math/rand"
	"time"
)

const maxRandomNumber = 100

var uniqIndex = 1

func InitArray(size int) []int {
	array := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		minus := -1
		if r.Intn(2) == 1 {
			minus = 1
		}
		array[i] = r.Intn(maxRandomNumber) * minus
	}

	return array
}

func InitArrayWithUniqElements(size int) []int {
	array := make([]int, size)
	alreadyInserted := make(map[int]bool)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		minus := -1
		if r.Intn(2) == 1 {
			minus = 1
		}
		newElement := r.Intn(maxRandomNumber*size*uniqIndex) * minus
		if alreadyInserted[newElement] {
			i--
			continue
		}
		alreadyInserted[newElement] = true
		array[i] = newElement
	}
	return array
}

func RandElementOfArray(array []int) int {
	return array[rand.Intn(len(array)-1)]
}
