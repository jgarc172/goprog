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

func isSorted(ints []int) (sorted bool) {
	sorted = false
	return
}
