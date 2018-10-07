// a program which prompts the user to enter integers and
// stores the integers in a sorted slice
package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 0, 3)
	var num int
	var err error

	fmt.Println("Please enter several integer values.  To stop, enter 'X'")
	for err != io.EOF {
		fmt.Print("Enter an integer: ")
		num, err = ReadInt()
		if err == nil {
			ints = AppendSort(ints, num)
			fmt.Println(ints)
		}
	}
}

// ReadInt reads an integer from standard input
// returns io.EOF if it reads "X"
func ReadInt() (num int, err error) {
	strValue := ""
	fmt.Scan(&strValue)
	if strings.ToUpper(strValue) == "X" {
		return 0, io.EOF
	}
	return strconv.Atoi(strValue)
}

// AppendSort appends the integer value to the array then returns the sorted array
func AppendSort(arr []int, val int) []int {
	arr = append(arr, val)
	sort.Ints(arr)
	return arr
}
