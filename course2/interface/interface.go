// Package interface provides a shell
// to prompt the user for a command and arguments
package main

import (
	"fmt"
)

var usage = "usage: newanimal name arg | query name arg | exit"

func main() {
	var cmd command
	var name, arg string
	fmt.Println(usage)
	for cmd != "exit" {
		fmt.Print("> ")
		n, _ := fmt.Scanln(&cmd, &name, &arg)
		if n < 3 {
			if cmd != "exit" {
				fmt.Println(usage)
			}
			continue
		}
		cmd.invoke(name, arg)
	}
}

type command string

func (c command) invoke(name, arg string) {
	switch c {
	case "query":
		fmt.Println(c, name, arg)
	case "newanimal":
		fmt.Println(c, name, arg)
	default:
		fmt.Println(usage)
	}
}
