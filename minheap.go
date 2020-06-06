// This package is a go implementation of a min-heap
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PrintTree prints out a binary tree
func PrintTree(t []int) {
	for i := range t {
		fmt.Println(t[i])
	}
}

// RandomTree builds a random binary tree of Nodes
// with a maximum size bounded by maxsize and each
// node having a value of at most maxvalue
func RandomTree(maxsize int, maxvalue int) []int {
	rand.Seed(time.Now().UnixNano())
	randomsize := rand.Int() % (maxsize + 1)
	tree := make([]int, randomsize)

	// Build a list of orphan nodes of random values
	for i := 0; i < len(tree); i++ {
		tree[i] = rand.Int() % (maxvalue + 1)
	}
	return tree
}

func heapSort(arr []int) {
	var temp int
	n := len(arr)
	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		temp = arr[0]
		arr[0] = arr[i]
		arr[i] = temp

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// SortTree sorts the binary tree into a min heap
func heapify(arr []int, n int, i int) {
	largest := i // Initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2
	var temp int

	// If left child is larger than root
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	// If right child is larger than largest so far
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	// If largest is not root
	if largest != i {
		// swap values
		temp = arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp

		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

func main() {
	tree := RandomTree(15, 50)
	fmt.Println("tree before sort")
	PrintTree(tree)
	fmt.Println("tree after sort")
	heapSort(tree)
	PrintTree(tree)
}
