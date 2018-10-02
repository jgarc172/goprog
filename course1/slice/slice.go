// a program which prompts the user to enter integers and
// stores the integers in a sorted slice
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 0)
	strValue := ""

	fmt.Println("Please enter several integer values.  To stop, enter 'X'")
	for {
		fmt.Print("Enter an integer: ")
		fmt.Scan(&strValue)
		if strings.ToUpper(strValue) == "X" {
			break
		}
		intValue, err := strconv.Atoi(strValue)
		if err == nil {
			ints = append(ints, intValue)
			sort.Ints(ints)
			fmt.Println(ints)
		}
	}
}
