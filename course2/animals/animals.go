// Package animals provides a shell to enter animal commands
package main

import (
	"fmt"
)

func main() {
	var animal, method string
	fmt.Println("Usage: animal method | exit")
	for animal != "exit" {
		fmt.Print("> ")
		n, _ := fmt.Scanln(&animal, &method)
		if n < 2 {
			continue
		}
		fmt.Println(animal, method)
	}
}
