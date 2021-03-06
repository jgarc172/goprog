// Package main reads names from a file, adds them to a list and prints the list
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := Prompt("File name")
	file, fErr := os.Open(fileName)
	if fErr != nil {
		fmt.Println("failed to open file", fileName)
		os.Exit(1)
	}
	defer file.Close()

	names := readNames(file)
	printNames(names)
}

type name struct {
	fname string
	lname string
}

// Prompt prints prompt, reads and returns the answer
func Prompt(prompt string) (answer string) {
	fmt.Print(prompt + ": ")
	fmt.Scanln(&answer)
	return
}

// readNames reads and returns a slice of names from r
func readNames(r io.Reader) []name {
	names := make([]name, 0, 3)
	var n name
	var err error
	for err != io.EOF {
		n, err = readName(r)
		if err == nil {
			names = append(names, n)
		}
	}
	return names
}

// readName reads and returns the next name
// otherwise returns an io.EOF
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

// printNames prints the first and last names in each name
func printNames(names []name) {
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
