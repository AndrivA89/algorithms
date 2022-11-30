package search

import (
	"fmt"
	"time"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

var countOperation int

func SearchByTree(array []int, target int) int {
	fmt.Printf("SEARCH BY BINARY TREE array size %v elements\n", len(array))
	countOperation = 0

	startFillingTree := time.Now()
	tree := fillTree(array)

	fmt.Printf("duration filling tree = %v ms (%v ns)\n",
		time.Since(startFillingTree).Milliseconds(),
		time.Since(startFillingTree).Nanoseconds())
	fmt.Printf("number of operations = %d\n", countOperation)

	countOperation = 0

	startSearch := time.Now()
	result, ok := tree.find(target)

	fmt.Printf("duration search = %v ms (%v ns)\n",
		time.Since(startSearch).Milliseconds(),
		time.Since(startSearch).Nanoseconds())
	fmt.Printf("number of operations = %d\n\n", countOperation)

	if !ok {
		fmt.Printf("element '%v' wasn't found", target)
	}

	return result.value
}

func fillTree(array []int) *TreeNode {
	if len(array) == 0 {
		fmt.Println("ERROR: array is empty")
		return &TreeNode{}
	}

	// Инициализируем дерево и первый элемент добавляем в корень
	tree := &TreeNode{
		value: array[0],
	}
	countOperation++

	for i, element := range array {
		// Пропускаем первый элемент - он будет добавлен в корень
		if i == 0 {
			continue
		}
		tree.insert(element)
	}

	return tree
}

func (t *TreeNode) insert(element int) {
	if element == t.value {
		fmt.Printf("WARNING: element '%v' has already been inserted\n", element)
		return
	}

	if element < t.value {
		if t.left == nil {
			t.left = &TreeNode{value: element}
			countOperation++
			return
		}
		t.left.insert(element)
		return
	}

	if element > t.value {
		if t.right == nil {
			t.right = &TreeNode{value: element}
			countOperation++
			return
		}
		t.right.insert(element)
		return
	}
}

func (t *TreeNode) find(element int) (*TreeNode, bool) {
	if t == nil {
		fmt.Println("ERROR: tree is empty")
		return nil, false
	}
	
	switch {
	case element == t.value:
		countOperation++
		return t, true
	case element < t.value:
		countOperation++
		return t.left.find(element)
	default:
		countOperation++
		return t.right.find(element)
	}
}
