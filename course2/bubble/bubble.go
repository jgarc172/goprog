// Package bubble prompts the user for a sequence of integers to be
// printed in ascending order
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter a sequence of integer values, separated by a space.")

	ints := ReadInts()
	BubbleSort(ints)
	for _, n := range ints {
		fmt.Print(n, " ")
	}
	if len(ints) > 0 {
		fmt.Println("\nThe sequence is sorted:", isSorted(ints))
	}
}

// ReadInts reads integers from stdin and returns them as a slice of integers
func ReadInts() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	strs := strings.Split(str, " ")
	ints := make([]int, 0, 10)
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, n)
		}
	}
	return ints
}

// BubbleSort sorts slice ints in ascending order
func BubbleSort(ints []int) {
	// ints is a slice from 0 to N
	N := len(ints) - 1
	// assumes ints is not sorted
	// and repeats until ints is sorted
	isSorted := false
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

// isSorted returns weather or not the slice ints is sorted
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
