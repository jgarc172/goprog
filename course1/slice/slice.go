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
		num, err = readInt()
		if err == nil {
			ints = appendSort(ints, num)
			fmt.Println(ints)
		}
	}
}
func readInt() (num int, err error) {
	strValue := ""
	fmt.Scan(&strValue)
	if strings.ToUpper(strValue) == "X" {
		return 0, io.EOF
	}
	return strconv.Atoi(strValue)
}
func appendSort(arr []int, val int) []int {
	arr = append(arr, val)
	sort.Ints(arr)
	return arr
}
