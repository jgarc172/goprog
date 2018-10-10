// Package main creates a map of keys and values for
// Name and Address, obtained from stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var attr, value string
	person := make(map[string]string)

	attr = "name"
	value = PromptFor(attr)
	person[attr] = value
	attr = "address"
	value = PromptFor(attr)
	person[attr] = value
	fmt.Println(person)
}

// PromptFor prompts for an attribute and returns the value entered from stdin
func PromptFor(attrib string) (value string) {
	prompt := "Enter " + attrib + ": "
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	if ok := scanner.Scan(); ok {
		value = scanner.Text()
	}
	return
}
