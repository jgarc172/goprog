package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	aStr := scanner.Text()

	aStr = strings.ToLower(aStr)
	hasI := strings.HasPrefix(aStr, "i")
	hasN := strings.HasSuffix(aStr, "n")
	hasA := strings.Contains(aStr, "a")

	if hasI && hasN && hasA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
