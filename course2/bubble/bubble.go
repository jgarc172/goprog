// Package bubble provides BubbleSort and Swap functions that sort a slice of integers
package main

import (
	"fmt"
)

func main() {
	ints := make([]int, 4)
	BubbleSort(ints)
	fmt.Println("Sorted?:", isSorted(ints))
	fmt.Println(ints)

	ints = []int{3, 1, 4, 2}
	BubbleSort(ints)
	fmt.Println("Sorted?:", isSorted(ints))
	fmt.Println(ints)

	ints = []int{3, 0, 4, 0}
	BubbleSort(ints)
	fmt.Println("Sorted?:", isSorted(ints))
	fmt.Println(ints)

}

// BubbleSort sorts slice ints in ascending order
func BubbleSort(ints []int) {
	// ints is a slice from 0 to N
	N := len(ints) - 1
	// assumes ints is not sorted
	isSorted := false
	// repeat until ints is sorted
	for !isSorted {
		// now assumes ints is sorted
		// if every pair of elements is in order
		isSorted = true
		for i := 0; i < N; i++ {
			if !inOrder(ints[i], ints[i+1]) {
				isSorted = false
				Swap(ints, i)
			}
		}
	}

}

// Swap reverses the values at locations i and i + 1
func Swap(ints []int, i int) {
	ints[i], ints[i+1] = ints[i+1], ints[i]
}

// inOrder returns true if a and b are in ascending order
func inOrder(a, b int) bool {
	return a <= b
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
