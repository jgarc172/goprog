// Package main creates a JSON encoding of a map of keys and values for
// Name and Address, obtained from stdin
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var attr, value string
	contact := make(map[string]string)

	attr = "name"
	value = PromptFor(attr)
	contact[attr] = value
	attr = "address"
	value = PromptFor(attr)
	contact[attr] = value

	b, err := json.Marshal(contact)
	if err == nil {
		fmt.Println(string(b))
	}
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
