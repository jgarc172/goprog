package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	const exit = "X"
	intSlice := make([]int64, 0, 3)
	for i := 0; true; i++ {
		var input string
		fmt.Print("Input Int64, 'X' to exit: ")
		fmt.Scan(&input)
		if input == exit {
			break
		}
		i64, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Printf("not int64 with err: %g, input again\n", err)
			i--
			continue
		}

		intSlice = append(intSlice, i64)
	}
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] < intSlice[j]
	})
	fmt.Println(intSlice)
}
