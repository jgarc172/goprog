// Package bubble provides a Sort function that sorts a slice of integers
package main

import (
	"fmt"
)

func main() {
	ints := make([]int, 4)
	Sort(ints)
	sorted := isSorted(ints)
	fmt.Println("Sorted?:", sorted)
	ints = []int{3, 1, 4, 2}
	Sort(ints)
	sorted = isSorted(ints)
	fmt.Println("Sorted?:", sorted)

}

// Sort sorts slice ints
func Sort(ints []int) {

}

// isSorted returns weather or not ints is sorted
func isSorted(ints []int) (sorted bool) {
	// ints is a slice with int values at locations
	// from 0 to N
	// ints is sorted if all pairs of adjacent
	// values are in order
	N := len(ints) - 1
	for i := 0; i < N-1; i++ {
		if !inOrder(ints[i], ints[i+1]) {
			return
		}
	}
	sorted = true
	return
}

func inOrder(a, b int) bool {
	return a <= b
}
