// Package sort prompts the user for a sequence of integers to be
// printed in ascending order
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
	for _, n := range ints {
		fmt.Print(n, " ")
	}
	fmt.Println()

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

// Sort sorts a slice of integers in ascending order
func Sort(ints []int) {
	length := float64(len(ints))
	size := int(math.Round(length / 4.5))
	var arr [4][]int
	for i := 0; i < 3; i++ {
		arr[i] = ints[i*size : (i+1)*size]
	}
	arr[3] = ints[3*size:]
	fmt.Println(arr)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(n int) {
			sort.Ints(arr[n])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(arr)

	arrA := make([]int, len(arr[0])+len(arr[1]))
	wg.Add(1)
	go MergeSort(arr[0], arr[1], arrA, &wg)
	arrB := make([]int, len(arr[2])+len(arr[3]))
	wg.Add(1)
	go MergeSort(arr[2], arr[3], arrB, &wg)
	wg.Wait()

	fmt.Println(arrA, arrB)

	arrC := make([]int, len(arrA)+len(arrB))
	wg.Add(1)
	go MergeSort(arrA, arrB, arrC, &wg)
	wg.Wait()

	fmt.Println(arrC)
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

// MergeChan merges two channels of sorted integers into one sorted channel chan3
func MergeChan(chan1, chan2, chan3 chan int) {

	for {
	}
}
