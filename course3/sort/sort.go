// Package sort prompts the user for a sequence of integers to be
// split into 4 slices, sorted, then combined into a sorted sequence
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Please enter a sequence of integer values, separated by a space.")
	ints := ReadInts()
	Sort(ints)
}

// ReadInts reads a sequence of integers from stdin
// and returns them as a slice of integers
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

// Sort sorts and prints a slice of integers in ascending order
// by splitting the slice in 4 goroutines then combined
func Sort(ints []int) {
	fmt.Println("Splitting array into 4 arrays")
	arrays := Split4(ints)
	fmt.Println("and sorting them")
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Sorting array: ", arrays[n])
			sort.Ints(arrays[n])
			wg.Done()
		}(i)
	}
	wg.Wait()

	finalArr := Merge4(arrays)
	fmt.Println("1 sorted array")
	fmt.Println(finalArr)
}

// Split4 splits slice ints into an array of 4 slices
func Split4(ints []int) (newArr [4][]int) {
	length := float64(len(ints))
	size := int(math.Round(length / 4.5))
	for i := 0; i < 3; i++ {
		newArr[i] = ints[i*size : (i+1)*size]
	}
	newArr[3] = ints[3*size:]
	return
}

// Merge4 merges 4 sorted slices of integers into one sorted slice
func Merge4(arrays [4][]int) []int {
	var wg sync.WaitGroup
	arrA := make([]int, len(arrays[0])+len(arrays[1]))
	wg.Add(1)
	go MergeSort(arrays[0], arrays[1], arrA, &wg)
	arrB := make([]int, len(arrays[2])+len(arrays[3]))
	wg.Add(1)
	go MergeSort(arrays[2], arrays[3], arrB, &wg)
	wg.Wait()

	arrC := make([]int, len(arrA)+len(arrB))
	wg.Add(1)
	go MergeSort(arrA, arrB, arrC, &wg)
	wg.Wait()
	return arrC
}

// MergeSort merges two sorted arrays into the third array
func MergeSort(arr1, arr2, arr3 []int, wg *sync.WaitGroup) {
	size1 := len(arr1)
	size2 := len(arr2)
	i := 0
	j := 0
	k := 0
	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			arr3[k] = arr1[i]
			i++
		} else {
			arr3[k] = arr2[j]
			j++
		}
		k++
	}

	for i < size1 {
		arr3[k] = arr1[i]
		i++
		k++
	}
	for j < size2 {
		arr3[k] = arr2[j]
		j++
		k++
	}
	wg.Done()
}
