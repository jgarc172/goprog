// Package read prints a list of names
package main

import (
	"fmt"
)

func main() {
	names := make([]name, 0, 3)
	n := name{"Joe", "Doe"}
	names = append(names, n)
	n = name{"John", "Blow"}
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
