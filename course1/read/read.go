// Package read repeatedly reads names from stdin, adds them to a list and prints the list
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	names := make([]name, 0, 3)
	var n name
	var err error
	for err != io.EOF {
		n, err = readName(os.Stdin)
		if err == nil {
			names = append(names, n)
		}
	}
	fmt.Println(names)
	printNames(names)
}

type name struct {
	fname string
	lname string
}

func readName(r io.Reader) (name, error) {
	var n name
	var f, l string
	_, err := fmt.Fscanln(r, &f, &l)
	if err != nil {
		return n, io.EOF
	}
	n = name{f, l}
	return n, nil
}
func printNames(names []name) {
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
