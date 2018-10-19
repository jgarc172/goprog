// Package interface provides a shell
// to prompt the user for a command and arguments
package main

import (
	"fmt"
)

func main() {
	var command, name, arg string
	usage := "usage: newanimal name arg | query name arg | exit"
	fmt.Println(usage)
	for command != "exit" {
		fmt.Print("> ")
		n, _ := fmt.Scanln(&command, &name, &arg)
		if n < 3 {
			if command != "exit" {
				fmt.Println(usage)
			}
			continue
		}
		fmt.Println(command)
	}
}
