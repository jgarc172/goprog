// Package main prompts the user to enter any string.
// Then the program prints either "Found" or "NotFound".
// Found: when the string starts with "i" contains "a" and ends with "n"
// Not Found: otherwise
// Search strings are not case sensitive.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	answer := Prompt("Enter a string")

	if Contains(answer, "i", "a", "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

// Prompt prints prompt, reads and returns the answer
func Prompt(prompt string) (answer string) {
	fmt.Print(prompt + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer = scanner.Text()
	return
}

// Contains returns whether all the corresponding elements are
// at the start, end, and in the string str
func Contains(str, start, in, end string) bool {
	str = strings.ToLower(str)

	starts := strings.HasPrefix(str, start)
	ends := strings.HasSuffix(str, end)
	contains := strings.Contains(str, in)

	return starts && ends && contains
}
