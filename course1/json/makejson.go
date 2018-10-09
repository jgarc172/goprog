// Package main prompts user for Name and Address
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name, address string
	name = PromptFor("Name")
	fmt.Println("name is", name)
	address = PromptFor("Address")
	fmt.Println("address is", address)
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
