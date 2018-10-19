// Package interface provides a shell
// to prompt the user for an animal name
package main

import (
	"fmt"
)

func main() {
	var name string
	usage := "Enter: animal-name | exit"
	fmt.Println(usage)
	for name != "exit" {
		fmt.Print("> ")
		fmt.Scanln(&name)
		if name != "exit" {
			fmt.Println(name)
		}
	}
}
