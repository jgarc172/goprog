// Package read reads names from stdin, adds them to a list and prints the list
package main

import (
	"fmt"
)

func main() {
	names := make([]name, 0, 3)
	var n name
	var f, l string
	fmt.Scanln(&f, &l)
	n = name{f, l}
	names = append(names, n)
	fmt.Scanln(&f, &l)
	n = name{f, l}
	names = append(names, n)
	fmt.Println(names)
	printNames(names)
}

type name struct {
	fname string
	lname string
}

func printNames(names []name) {
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
